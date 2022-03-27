package uniclass_ol_ea_com_exporters

import (
	"fmt"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com"
	domain_to_ol_ea_com_migration_orchestrator "github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/domain_migration/domain_to_ol_ea_com_migration/orchestrators"
	ea_tools_session_orchestrator "github.com/OntoLedgy/ol_ea_common_tools/code/services/session/orchestrators"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
	"path/filepath"
)

//import os
//from nf_common_source.code.services.file_system_service.objects.folders import Folders
//from nf_common_source.code.services.reporting_service.reporters.log_with_datetime import log_message
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.nf_ea_com_universes import NfEaComUniverses
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.domain_migration.domain_to_nf_ea_com_migration.orchestrators.nf_ea_com_universe_to_eapx_migration_orchestator import \
//orchestrate_nf_ea_com_universe_to_eapx_migration
//from nf_ea_common_tools_source.b_code.services.session.orchestrators.ea_tools_session_managers import \
//EaToolsSessionManagers

//def export_nf_ea_com_to_ea(
func ExportNfEaComToEa(
	//ea_tools_session_manager: EaToolsSessionManagers,
	ea_tools_session_manager ea_tools_session_orchestrator.EaToolsSessionManagers,
	//nf_ea_com_universe: NfEaComUniverses,
	nf_ea_com_universe *com.OlEaComUniverses,
	//folder_path: str):
	folder_path string) {

	//bclearer_stage = \
	bclearer_stage :=
		//nf_ea_com_universe.ea_repository.short_name
		nf_ea_com_universe.EaRepositories.ShortName

	outputFolder := &object_model.Folders{}

	//current_stage_ea_export_folder_path = \
	current_stage_ea_export_folder_path :=

		//os.path.join(
		filepath.Join(
			//folder_path,
			folder_path,
			//bclearer_stage,
			bclearer_stage,
			//bclearer_stage + '_ea_export')
			bclearer_stage+"_ea_export")

	outputFolder.Initialise(
		current_stage_ea_export_folder_path,
		nil)

	outputFolder.CreateIfNonExistent()

	//if not os.path.exists(current_stage_ea_export_folder_path):
	//os.mkdir(
	//current_stage_ea_export_folder_path)

	//log_message(
	//'STARTING EA EXPORT (XML) FOR ' + bclearer_stage)
	fmt.Println(
		"STARTING EA EXPORT (XML) FOR " + bclearer_stage)

	//orchestrate_nf_ea_com_universe_to_eapx_migration(
	domain_to_ol_ea_com_migration_orchestrator.OrchestrateOlEaComUniverseToEapxMigration(
		//ea_tools_session_manager=ea_tools_session_manager,
		ea_tools_session_manager,
		//nf_ea_com_universe=nf_ea_com_universe,
		nf_ea_com_universe,
		//short_name=bclearer_stage,
		bclearer_stage,
		//output_folder=Folders(
		//absolute_path_string=folder_path))
		outputFolder)

	//log_message(
	//'EA EXPORT DONE (XML) - ' + bclearer_stage)
}
