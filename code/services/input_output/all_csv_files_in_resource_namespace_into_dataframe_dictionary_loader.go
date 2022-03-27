package input_output

//https://github.com/boro-alpha/uniclass_to_nf_ea_com/blob/master/uniclass_to_nf_ea_com_source/b_code/services/input_output/all_csv_files_in_resource_namespace_into_dataframe_dictionary_loader.py

//import importlib
//from os.path import splitext, basename
//import pandas
//from nf_common_source.code.services.file_system_service.files_of_extension_from_folder_getter import \
//get_all_files_of_extension_from_folder
//from nf_common_source.code.services.file_system_service.objects.files import Files
//from nf_common_source.code.services.file_system_service.objects.folders import Folders
//from uniclass_to_nf_ea_com_source.b_code.configurations.common_constants.uniclass_bclearer_constants import CSV_EXTENSION_FILE_NAME

import (
	"bclearer_projects/uniclass/code/configurations/common_constants"
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"
)

//def load_all_csv_files_in_resource_namespace_into_dataframe_dictionary(
func LoadAllCsvFilesInResourceNamespaceIntoDataframeMap(
	//resource_namespace: str) \
	//-> dict:
	resource_namespace string) map[string]*dataframes.DataFrames {

	//dictionary_of_dataframes = \
	//dict()
	mapOfDataframes :=
		map[string]*dataframes.DataFrames{}

	//TODO - investigate how importing module in go would work.
	//module = \
	//importlib.import_module(
	//name=resource_namespace)

	//module_path_string = \
	//module.__path__[0]

	//resource_namespace_folder = \
	resource_namespace_folder :=
		//Folders(
		&object_model.Folders{}
	//absolute_path_string=module_path_string)

	resource_namespace_folder.Initialise(
		"D:\\S\\go\\src\\bclearer_projects\\uniclass\\resources\\1_load_inputs\\1l_s3_uniclass_tables\\full_scope",
		nil)

	//csv_files = \
	csvFiles :=
		//get_all_files_of_extension_from_folder(
		file_system_service.GetAllFilesOfExtensionFromFolder(
			//folder=resource_namespace_folder,
			resource_namespace_folder,
			//dot_extension_string=CSV_EXTENSION_FILE_NAME)
			common_constants.CSV_EXTENSION_FILE_NAME)

	//for csv_file in csv_files:
	for csvFile :=
		csvFiles.Front(); csvFile != nil; csvFile = csvFile.Next() {

		//dictionary_of_dataframes = \
		mapOfDataframes =
			//__add_csv_file_to_dictionary_of_dataframes(
			addCsvFileToDictionaryOfDataframes(
				//csv_file=csv_file,
				csvFile.Value.(*object_model.Files),
				//dictionary_of_dataframes=dictionary_of_dataframes)
				mapOfDataframes)
		fmt.Println(csvFile.Value.(*object_model.Files).Path.String())

	}

	//return \
	//dictionary_of_dataframes
	return mapOfDataframes
}

//def __add_csv_file_to_dictionary_of_dataframes(
func addCsvFileToDictionaryOfDataframes(
	//csv_file: Files,
	csv_file *object_model.Files,
	//dictionary_of_dataframes: dict) \
	//-> dict:
	MapOfDataframes map[string]*dataframes.DataFrames) map[string]*dataframes.DataFrames {

	//base_file_name_without_extension = \

	//splitext(
	//basename(
	//csv_file.absolute_path_string))[0]

	//dataframe = \
	//pandas.read_csv(
	//filepath_or_buffer=csv_file.absolute_path_string,
	//dtype=str,
	//encoding='cp1252')

	//dictionary_of_dataframes.update({
	//base_file_name_without_extension: dataframe})

	//return \
	//dictionary_of_dataframes
	return MapOfDataframes
}
