package main

import (
	"bclearer_projects/uniclass/code/configurations/resource_constants"
	"bclearer_projects/uniclass/code/orchestrators"
)

//from uniclass_to_nf_ea_com_source.b_code.configurations.resource_constants.resources_namespace_constants import \
//LOAD_4_INPUT_TABLES_RESOURCES_NAMESPACE
//from uniclass_to_nf_ea_com_source.b_code.orchestrators.uniclass_bclearer_orchestrator import \
//orchestrate_uniclass_bclearer

//if __name__ == '__main__':
func main() {

	//orchestrate_uniclass_bclearer(
	orchestrators.OrchestrateUniclassbClearer(
		//uniclass_source_data_resource_namespace=LOAD_4_INPUT_TABLES_RESOURCES_NAMESPACE)
		resource_constants.LOAD_4_INPUT_TABLES_RESOURCES_NAMESPACE)

}
