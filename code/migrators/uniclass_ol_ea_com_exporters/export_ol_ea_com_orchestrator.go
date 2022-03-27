package uniclass_ol_ea_com_exporters

//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.temporary.nf_ea_com_dictionary_to_nf_ea_com_universe_convertor import \
//convert_nf_ea_com_dictionary_to_nf_ea_com_universe
//from nf_ea_common_tools_source.b_code.services.session.orchestrators.ea_tools_session_managers import \
//EaToolsSessionManagers
//from uniclass_to_nf_ea_com_source.b_code.migrators.uniclass_nf_ea_com_exporters.nf_ea_com_to_access_exporter import \
//export_nf_ea_com_to_access
//from uniclass_to_nf_ea_com_source.b_code.migrators.uniclass_nf_ea_com_exporters.nf_ea_com_to_ea_exporter import \
//export_nf_ea_com_to_ea
//from uniclass_to_nf_ea_com_source.b_code.migrators.uniclass_nf_ea_com_exporters.nf_ea_com_to_hdf5_exporter import \
//export_nf_ea_com_to_hdf5

import (
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com/common_knowledge/collection_types"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com/temporary"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/orchestrators"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"
)

//def orchestrate_export_nf_ea_com(
func OrchestrateExportOlEaCom(
	//nf_ea_com_dictionary: dict,
	olEaComMap map[collection_types.OlEaComCollectionTypes]dataframes.DataFrames,
	//output_folder_path: str,
	outputFolderPath string,
	//bclearer_stage: str)
	bclearerStage string) {

	var eaToolSessionManager = orchestrators.EaToolsSessionManagers{}
	//with EaToolsSessionManagers() as ea_tools_session_manager:

	//nf_ea_com_universe = \
	ol_ea_com_universe :=
		//convert_nf_ea_com_dictionary_to_nf_ea_com_universe(
		temporary.ConvertOlEaComMapToOlEaComUniverse(
			//ea_tools_session_manager=ea_tools_session_manager,
			eaToolSessionManager,
			//nf_ea_com_dictionary=nf_ea_com_dictionary,
			olEaComMap,
			//short_name=bclearer_stage)
			bclearerStage)

	//export_nf_ea_com_to_access(
	//ea_tools_session_manager=ea_tools_session_manager,
	//nf_ea_com_universe=nf_ea_com_universe,
	//folder_path=output_folder_path)

	//export_nf_ea_com_to_ea(
	ExportNfEaComToEa(
		//ea_tools_session_manager=ea_tools_session_manager,
		eaToolSessionManager,
		//nf_ea_com_universe=nf_ea_com_universe,
		ol_ea_com_universe,
		//folder_path=output_folder_path)
		outputFolderPath)

	//# export_nf_ea_com_to_hdf5(
	//#     nf_ea_com_universe=nf_ea_com_universe,
	//#     folder_path=output_folder_path)

}
