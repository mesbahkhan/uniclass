package orchestrators

//https://github.com/boro-alpha/uniclass_to_nf_ea_com/blob/master/uniclass_to_nf_ea_com_source/b_code/orchestrators/uniclass_bclearer_orchestrator.py

//import os
//from nf_common_source.code.services.datetime_service.time_helpers.time_getter import now_time_as_string_for_files
//from nf_common_source.code.services.file_system_service.folder_selector import select_folder
//from uniclass_to_nf_ea_com_source.b_code.orchestrators.stages.evolve.evolve_stage_1_orchestrator import orchestrate_evolve_stage_1
//from uniclass_to_nf_ea_com_source.b_code.orchestrators.stages.evolve.evolve_stage_2_orchestrator import orchestrate_evolve_stage_2
//from uniclass_to_nf_ea_com_source.b_code.orchestrators.stages.evolve.evolve_stage_3_orchestrator import orchestrate_evolve_stage_3
//from uniclass_to_nf_ea_com_source.b_code.orchestrators.stages.evolve.evolve_stage_4_orchestrator import orchestrate_evolve_stage_4
//from uniclass_to_nf_ea_com_source.b_code.orchestrators.stages.evolve.evolve_stage_5_orchestrator import orchestrate_evolve_stage_5
//from uniclass_to_nf_ea_com_source.b_code.orchestrators.stages.evolve.evolve_stage_7_orchestrator import orchestrate_evolve_stage_7
//from uniclass_to_nf_ea_com_source.b_code.orchestrators.stages.evolve.evolve_stage_8_orchestrator import orchestrate_evolve_stage_8
//from uniclass_to_nf_ea_com_source.b_code.orchestrators.stages.evolve.evolve_stage_6_orchestrator import orchestrate_evolve_stage_6
//from uniclass_to_nf_ea_com_source.b_code.orchestrators.stages.load.load_stage_4_orchestrator import orchestrate_load_stage_4
//from uniclass_to_nf_ea_com_source.b_code.configurations.common_constants.uniclass_bclearer_constants import ROOT_OUTPUT_FOLDER_TITLE_MESSAGE
//from uniclass_to_nf_ea_com_source.b_code.configurations.common_constants.uniclass_bclearer_constants import UNICLASS_BCLEARER_PREFIX_FOR_NAMES
//from uniclass_to_nf_ea_com_source.b_code.services.input_output.set_up_and_close_out_logging import set_up_logger_and_output_folder
//from uniclass_to_nf_ea_com_source.b_code.services.input_output.set_up_and_close_out_logging import close_log_file

import (
	"bclearer_projects/uniclass/code/configurations/common_constants"
	"bclearer_projects/uniclass/code/orchestrators/stages/load"
	"github.com/OntoLedgy/ol_common_services/code/services/datetime_service/time_helpers"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"
	"os"
	"strings"
)

//def orchestrate_uniclass_bclearer(
func OrchestrateUniclassbClearer(
	//uniclass_source_data_resource_namespace: str):
	uniclassSourceDataResourceNamespace string) {

	//root_folder_path = \
	rootFolderPath :=
		//__get_output_root_folder_path()
		getOutputRootFolderPath()

	//TODO --> add logging infra
	//set_up_logger_and_output_folder(
	//-->SetUpLoggerAndOutputFolder(
	//output_folder_name=root_folder_path)
	//-->root_folder_path)

	//load_stage_4_dictionary_of_dataframes = \
	loadStage4MapOfDataframes :=
		//__orchestrate_uniclass_load_stage(
		orchestrateUniclassLoadStage(
			//folder_path=root_folder_path,
			rootFolderPath,
			//uniclass_source_data_resource_namespace=uniclass_source_data_resource_namespace)
			uniclassSourceDataResourceNamespace)

	//__orchestrate_uniclass_evolve_stage(
	orchestrateUniclassEvolveStage(
		//load_stage_4_dictionary_of_dataframes=load_stage_4_dictionary_of_dataframes,
		loadStage4MapOfDataframes,
		//folder_path=root_folder_path)
		rootFolderPath)

	//TODO --> add logging infra
	//close_log_file()
	//--> CloseLogFile()

}

//def __get_output_root_folder_path() \
//-> str:
func getOutputRootFolderPath() string {

	//root_folder = \
	rootFolder :=
		//select_folder(
		file_system_service.SelectFolder(common_constants.ROOT_OUTPUT_FOLDER_TITLE_MESSAGE)
	//title=ROOT_OUTPUT_FOLDER_TITLE_MESSAGE)

	//root_folder_path = \
	rootFolderPath :=
		rootFolder.AbsolutePathString()
	//root_folder.absolute_path_string

	//output_root_folder_path = \
	output_root_folder_path :=
		rootFolder.Path.Join(
			//os.path.join(
			//root_folder_path,
			common_constants.UNICLASS_BCLEARER_PREFIX_FOR_NAMES +
				strings.ReplaceAll(time_helpers.NowTimeAsStringForFiles(), "_", ""))
	//UNICLASS_BCLEARER_PREFIX_FOR_NAMES + now_time_as_string_for_files().replace('_', ''))

	//os.mkdir(
	os.Mkdir(
		//output_root_folder_path)
		output_root_folder_path.String(), 0755)

	//return \
	//output_root_folder_path
	return rootFolderPath
}

//def __orchestrate_uniclass_load_stage(
func orchestrateUniclassLoadStage(
	//folder_path: str,
	folderPath string,
	//uniclass_source_data_resource_namespace: str) \
	//-> dict:
	uniclassSourceDataResourceNamespace string) map[string]*dataframes.DataFrames {

	//load_stage_4_dictionary_of_dataframes = \
	stage4MapOfDataframes :=
		//orchestrate_load_stage_4(
		load.OrchestrateLoadStage4(
			//folder_path=folder_path,
			folderPath,
			//uniclass_source_data_resource_namespace=uniclass_source_data_resource_namespace)
			uniclassSourceDataResourceNamespace)

	//return \
	//load_stage_4_dictionary_of_dataframes
	return stage4MapOfDataframes
}

//def __orchestrate_uniclass_evolve_stage(
func orchestrateUniclassEvolveStage(
	//load_stage_4_dictionary_of_dataframes: dict,
	load_stage_4_dictionary_of_dataframes map[string]*dataframes.DataFrames,
	//folder_path: str):
	folder_path string) {

	//evolve_stage_1_dictionary_of_dataframes = \
	//orchestrate_evolve_stage_1(
	//dictionary_of_dataframes=load_stage_4_dictionary_of_dataframes,
	//folder_path=folder_path)

	//evolve_stage_2_dictionary_of_dataframes = \
	//orchestrate_evolve_stage_2(
	//dictionary_of_dataframes=evolve_stage_1_dictionary_of_dataframes,
	//folder_path=folder_path)

	//evolve_stage_3_dictionary_of_dataframes = \
	//orchestrate_evolve_stage_3(
	//dictionary_of_dataframes=evolve_stage_2_dictionary_of_dataframes,
	//folder_path=folder_path)

	//evolve_stage_4_dictionary_of_dataframes = \
	//orchestrate_evolve_stage_4(
	//dictionary_of_dataframes=evolve_stage_3_dictionary_of_dataframes,
	//folder_path=folder_path)

	//evolve_stage_5_dictionary_of_dataframes = \
	//orchestrate_evolve_stage_5(
	//dictionary_of_dataframes=evolve_stage_4_dictionary_of_dataframes,
	//folder_path=folder_path)

	//evolve_stage_6_dictionary_of_dataframes = \
	//orchestrate_evolve_stage_6(
	//dictionary_of_dataframes=evolve_stage_5_dictionary_of_dataframes,
	//folder_path=folder_path)

	//evolve_stage_7_dictionary_of_dataframes = \
	//orchestrate_evolve_stage_7(
	//dictionary_of_dataframes=evolve_stage_6_dictionary_of_dataframes,
	//folder_path=folder_path)

	//evolve_stage_8_dictionary_of_dataframes = \
	//orchestrate_evolve_stage_8(
	//evolve_stage_7_dictionary_of_dataframes=evolve_stage_7_dictionary_of_dataframes,
	//folder_path=folder_path

}
