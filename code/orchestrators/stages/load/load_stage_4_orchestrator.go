package load

//from nf_common_source.code.services.reporting_service.wrappers.run_and_log_function_wrapper import run_and_log_function

//from uniclass_to_nf_ea_com_source.b_code.migrators.uniclass_domain_to_nf_ea_com.load.load_stage_4.nf_ea_com_tables_creation_for_load_stage_4_orchestrator import \
//orchestrate_nf_ea_com_tables_creation_for_load_stage_4
//from uniclass_to_nf_ea_com_source.b_code.migrators.uniclass_nf_ea_com_exporters.export_nf_ea_com_orchestrator import \
//orchestrate_export_nf_ea_com
//from uniclass_to_nf_ea_com_source.b_code.migrators.uniclass_raw_to_domain.load.load_stage_4.load_stage_4_domain_tables_generation_orchestrator import \
//orchestrate_domain_tables_creation_for_load_stage_4

import (
	load_stage_42 "bclearer_projects/uniclass/code/migrators/uniclass_domain_to_ol_ea_com/load/load_stage_4"
	"bclearer_projects/uniclass/code/migrators/uniclass_ol_ea_com_exporters"
	"bclearer_projects/uniclass/code/migrators/uniclass_raw_to_domain/load/load_stage_4"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"
)

//@run_and_log_function
//def orchestrate_load_stage_4(
func OrchestrateLoadStage4(
	//folder_path: str,
	folderPath string,
	//uniclass_source_data_resource_namespace: str)\
	//-> dict:
	uniclassSourceDataResourceNamespace string) map[string]*dataframes.DataFrames {

	//uniclass2015_uuidified_dictionary_of_dataframes = \
	uniclass2015UuidifiedMapOfDataframes :=
		//orchestrate_domain_tables_creation_for_load_stage_4(
		load_stage_4.OrchestrateDomainTablesCreationForLoadStage4(
			//folder_path=folder_path,
			folderPath,
			//uniclass_source_data_resource_namespace=uniclass_source_data_resource_namespace)
			uniclassSourceDataResourceNamespace)

	//load_stage_4_nf_ea_com_tables = \
	loadStage4NfEaComTables :=
		//orchestrate_nf_ea_com_tables_creation_for_load_stage_4(
		load_stage_42.OrchestrateNfEaComTablesCreationForLoadStage4(
			//dictionary_of_dataframes=uniclass2015_uuidified_dictionary_of_dataframes)
			uniclass2015UuidifiedMapOfDataframes)

	//orchestrate_export_nf_ea_com(
	uniclass_ol_ea_com_exporters.OrchestrateExportOlEaCom(
		//nf_ea_com_dictionary=load_stage_4_nf_ea_com_tables,
		loadStage4NfEaComTables,
		//output_folder_path=folder_path,
		folderPath,
		//bclearer_stage='load_4')
		"load_4")

	//return \
	//uniclass2015_uuidified_dictionary_of_dataframes
	return uniclass2015UuidifiedMapOfDataframes
}
