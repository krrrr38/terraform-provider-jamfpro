// computerinventory_data_source.go
package computerinventory

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceJamfProMacOSComputerInventory provides information about a specific computer's inventory by its ID or Name.
func DataSourceJamfProComputerInventory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfProComputerInventoryRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"udid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"general": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_reported_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"jamf_binary_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"barcode1": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"barcode2": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"asset_tag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_management": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"managed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"management_username": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"supervised": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mdm_capable": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"capable": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"capable_users": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"report_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_contact_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_cloud_backup_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_enrolled_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mdm_profile_expiration": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"initial_entry_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribution_point": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enrollment_method": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"site": {
							Type:     schema.TypeList,
							Computed: true,

							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"itunes_store_account_active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enrolled_via_automated_device_enrollment": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"user_approved_mdm": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"declarative_device_management_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"extension_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"definition_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"multi_value": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"values": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"data_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"input_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"management_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"disk_encryption": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"boot_partition_encryption_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"partition_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"partition_file_vault2_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"partition_file_vault2_percent": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"individual_recovery_key_validity_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"institutional_recovery_key_present": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"disk_encryption_configuration_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"file_vault2_enabled_user_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"file_vault2_eligibility_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"purchasing": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"leased": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"purchased": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"po_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"po_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vendor": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"warranty_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"apple_care_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lease_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"purchase_price": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"life_expectancy": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"purchasing_account": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"purchasing_contact": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"extension_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"definition_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"multi_value": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"values": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"data_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"input_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"applications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_app_store": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"size_megabytes": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bundle_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_available": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"external_version_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"storage": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"boot_drive_available_space_megabytes": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"disks": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"device": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"model": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"revision": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"serial_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size_megabytes": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"smart_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"partitions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"size_megabytes": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"available_megabytes": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"partition_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"percent_used": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"file_vault2_state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"file_vault2_progress_percent": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"lvm_managed": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"user_and_location": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"realname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"position": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"department_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"building_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"room": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"extension_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"definition_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"multi_value": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"values": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"data_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"input_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"configuration_profiles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_installed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"removable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"printers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"hardware": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"make": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model_identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"processor_speed_mhz": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"processor_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"processor_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"processor_architecture": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bus_speed_mhz": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cache_size_kilobytes": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_adapter_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"alt_network_adapter_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"alt_mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_ram_megabytes": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"open_ram_slots": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"battery_capacity_percent": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"smc_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nic_speed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"optical_drive": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"boot_rom": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ble_capable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"supports_ios_app_installs": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"apple_silicon": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"extension_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"definition_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"multi_value": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"values": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"data_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"input_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"local_user_accounts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_guid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"full_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"admin": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"home_directory": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"home_directory_size_mb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"file_vault2_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"user_account_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_min_length": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"password_max_age": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"password_min_complex_characters": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"password_history_depth": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"password_require_alphanumeric": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"computer_azure_active_directory_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_azure_active_directory_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"azure_active_directory_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"certificates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"common_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"expiration_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subject_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sha1_fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"issued_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"file_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size_bytes": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"plugins": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"package_receipts": {
				Type:     schema.TypeList,
				Computed: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"installed_by_jamf_pro": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"installed_by_installer_swu": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"cached": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"fonts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"security": {
				Type:     schema.TypeList,
				Computed: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sip_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gatekeeper_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"xprotect_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_login_disabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"remote_desktop_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"activation_lock_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"recovery_lock_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"firewall_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"secure_boot_level": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_boot_level": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bootstrap_token_allowed": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"operating_system": {
				Type:     schema.TypeList,
				Computed: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"build": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supplemental_build_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rapid_security_response": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"active_directory_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"file_vault2_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_update_device_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"extension_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"definition_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"multi_value": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"values": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"data_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"input_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"licensed_software": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ibeacons": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"software_updates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"package_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"extension_attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"definition_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"multi_value": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"values": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"data_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"options": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"input_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			// ignoring contentCashing section intentionally as it doesnt serve any use in terraform scenarios
			"group_memberships": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"smart_group": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// dataSourceJamfProComputerInventoryRead fetches the details of a specific macOS Configuration Profile
// from Jamf Pro using either its unique Name or its ID. The function prioritizes the 'name' attribute over the 'id'
// attribute for fetching details. If neither 'name' nor 'id' is provided, it returns an error.
// Once the details are fetched, they are set in the data source's state.
func dataSourceJamfProComputerInventoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Asserts 'meta' as '*client.APIClient'
	apiclient, ok := meta.(*client.APIClient)
	if !ok {
		return diag.Errorf("error asserting meta as *client.APIClient")
	}
	conn := apiclient.Conn

	var profile *jamfpro.ResponseComputerInventory
	var err error

	// Fetch profile by 'name' or 'id'
	if v, ok := d.GetOk("name"); ok {
		profileName, ok := v.(string)
		if !ok {
			return diag.Errorf("error asserting 'name' as string")
		}
		profile, err = conn.GetComputerInventoryByName(profileName)
	} else if v, ok := d.GetOk("id"); ok {
		profileID, ok := v.(string)
		if !ok {
			return diag.Errorf("error asserting 'id' as string")
		}
		profile, err = conn.GetComputerInventoryByID(profileID)
	} else {
		return diag.Errorf("Either 'name' or 'id' must be provided")
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch macOS Configuration Profile: %v", err))
	}

	// Set top-level attributes
	d.SetId(profile.ID)
	d.Set("id", profile.ID)
	d.Set("udid", profile.UDID)

	// Set 'general' section
	if err := setGeneralSection(d, profile.General); err != nil {
		return diag.FromErr(err)
	}

	// Set 'diskEncryption' section
	if err := setDiskEncryptionSection(d, profile.DiskEncryption); err != nil {
		return diag.FromErr(err)
	}

	// Set 'purchasing' section
	if err := setPurchasingSection(d, profile.Purchasing); err != nil {
		return diag.FromErr(err)
	}

	// Set 'applications' section
	if err := setApplicationsSection(d, profile.Applications); err != nil {
		return diag.FromErr(err)
	}

	// Set 'storage' section
	if err := setStorageSection(d, profile.Storage); err != nil {
		return diag.FromErr(err)
	}

	// Set 'userAndLocation' section
	if err := setUserAndLocationSection(d, profile.UserAndLocation); err != nil {
		return diag.FromErr(err)
	}

	// Set 'hardware' section
	if err := setHardwareSection(d, profile.Hardware); err != nil {
		return diag.FromErr(err)
	}

	// Set 'hardware' section
	if err := setHardwareSection(d, profile.Hardware); err != nil {
		return diag.FromErr(err)
	}

	// Set 'localUserAccounts' section
	if err := setLocalUserAccountsSection(d, profile.LocalUserAccounts); err != nil {
		return diag.FromErr(err)
	}

	// Set 'certificates' section
	if err := setCertificatesSection(d, profile.Certificates); err != nil {
		return diag.FromErr(err)
	}

	// Set 'attachments' section
	if err := setAttachmentsSection(d, profile.Attachments); err != nil {
		return diag.FromErr(err)
	}

	// Set 'plugins' section
	if err := setPluginsSection(d, profile.Plugins); err != nil {
		return diag.FromErr(err)
	}

	// Set 'packageReceipts' section
	if err := setPackageReceiptsSection(d, profile.PackageReceipts); err != nil {
		return diag.FromErr(err)
	}

	// Set 'fonts' section
	if err := setFontsSection(d, profile.Fonts); err != nil {
		return diag.FromErr(err)
	}

	// Set 'security' section
	if err := setSecuritySection(d, profile.Security); err != nil {
		return diag.FromErr(err)
	}

	// Set 'operatingSystem' section
	if err := setOperatingSystemSection(d, profile.OperatingSystem); err != nil {
		return diag.FromErr(err)
	}

	// Set 'licensedSoftware' section
	if err := setLicensedSoftwareSection(d, profile.LicensedSoftware); err != nil {
		return diag.FromErr(err)
	}

	// Set 'ibeacons' section
	if err := setIBeaconsSection(d, profile.Ibeacons); err != nil {
		return diag.FromErr(err)
	}

	// Set 'softwareUpdates' section
	if err := setSoftwareUpdatesSection(d, profile.SoftwareUpdates); err != nil {
		return diag.FromErr(err)
	}

	// Set 'extensionAttributes' section
	if err := setExtensionAttributesSection(d, profile.ExtensionAttributes); err != nil {
		return diag.FromErr(err)
	}

	// Set 'groupMemberships' section
	if err := setGroupMembershipsSection(d, profile.GroupMemberships); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// setGeneralSection maps the 'general' section of the computer inventory response to the Terraform schema.
func setGeneralSection(d *schema.ResourceData, general jamfpro.ComputerInventoryDataSubsetGeneral) error {
	// Initialize a map to hold the 'general' section attributes.
	gen := make(map[string]interface{})

	// Map each attribute of the 'general' section from the API response to the corresponding Terraform schema attribute.
	gen["name"] = general.Name
	gen["last_Ip_Address"] = general.LastIpAddress
	gen["last_Reported_Ip"] = general.LastReportedIp
	gen["jamfBinaryVersion"] = general.JamfBinaryVersion
	gen["platform"] = general.Platform
	gen["barcode1"] = general.Barcode1
	gen["barcode2"] = general.Barcode2
	gen["assetTag"] = general.AssetTag
	gen["supervised"] = general.Supervised
	gen["mdmCapable"] = general.MdmCapable.Capable
	gen["reportDate"] = general.ReportDate
	gen["lastContactTime"] = general.LastContactTime
	gen["lastCloudBackupDate"] = general.LastCloudBackupDate
	gen["lastEnrolledDate"] = general.LastEnrolledDate
	gen["mdmProfileExpiration"] = general.MdmProfileExpiration
	gen["initialEntryDate"] = general.InitialEntryDate
	gen["distributionPoint"] = general.DistributionPoint
	gen["itunesStoreAccountActive"] = general.ItunesStoreAccountActive
	gen["enrolledViaAutomatedDeviceEnrollment"] = general.EnrolledViaAutomatedDeviceEnrollment
	gen["userApprovedMdm"] = general.UserApprovedMdm
	gen["declarativeDeviceManagementEnabled"] = general.DeclarativeDeviceManagementEnabled
	gen["managementId"] = general.ManagementId

	// Handle nested object 'remoteManagement'.
	remoteManagement := make(map[string]interface{})
	remoteManagement["managed"] = general.RemoteManagement.Managed
	remoteManagement["managementUsername"] = general.RemoteManagement.ManagementUsername
	gen["remoteManagement"] = []interface{}{remoteManagement}

	// Handle nested object 'site'.
	if general.Site.ID != "" || general.Site.Name != "" {
		site := make(map[string]interface{})
		site["id"] = general.Site.ID
		site["name"] = general.Site.Name
		gen["site"] = []interface{}{site}
	}

	// Handle nested object 'enrollmentMethod'.
	if general.EnrollmentMethod.ID != "" || general.EnrollmentMethod.ObjectName != "" || general.EnrollmentMethod.ObjectType != "" {
		enrollmentMethod := make(map[string]interface{})
		enrollmentMethod["id"] = general.EnrollmentMethod.ID
		enrollmentMethod["objectName"] = general.EnrollmentMethod.ObjectName
		enrollmentMethod["objectType"] = general.EnrollmentMethod.ObjectType
		gen["enrollmentMethod"] = []interface{}{enrollmentMethod}
	}

	// Set the 'general' section in the Terraform resource data.
	return d.Set("general", []interface{}{gen})
}

// setDiskEncryptionSection maps the 'diskEncryption' section of the computer inventory response to the Terraform schema.
func setDiskEncryptionSection(d *schema.ResourceData, diskEncryption jamfpro.ComputerInventoryDataSubsetDiskEncryption) error {
	// Initialize a map to hold the 'diskEncryption' section attributes.
	diskEnc := make(map[string]interface{})

	// Map each attribute of the 'diskEncryption' section from the API response to the corresponding Terraform schema attribute.
	diskEnc["individualRecoveryKeyValidityStatus"] = diskEncryption.IndividualRecoveryKeyValidityStatus
	diskEnc["institutionalRecoveryKeyPresent"] = diskEncryption.InstitutionalRecoveryKeyPresent
	diskEnc["diskEncryptionConfigurationName"] = diskEncryption.DiskEncryptionConfigurationName
	diskEnc["fileVault2EligibilityMessage"] = diskEncryption.FileVault2EligibilityMessage

	// Handle nested object 'bootPartitionEncryptionDetails'.
	bootPartitionDetails := make(map[string]interface{})
	bootPartitionDetails["partitionName"] = diskEncryption.BootPartitionEncryptionDetails.PartitionName
	bootPartitionDetails["partitionFileVault2State"] = diskEncryption.BootPartitionEncryptionDetails.PartitionFileVault2State
	bootPartitionDetails["partitionFileVault2Percent"] = diskEncryption.BootPartitionEncryptionDetails.PartitionFileVault2Percent
	diskEnc["bootPartitionEncryptionDetails"] = []interface{}{bootPartitionDetails}

	// Map 'fileVault2EnabledUserNames' as a list of strings.
	fileVaultUserNames := make([]string, len(diskEncryption.FileVault2EnabledUserNames))
	copy(fileVaultUserNames, diskEncryption.FileVault2EnabledUserNames)

	// Set 'fileVault2EnabledUserNames' in the 'diskEnc' map.
	diskEnc["fileVault2EnabledUserNames"] = fileVaultUserNames

	// Set the 'diskEncryption' section in the Terraform resource data.
	return d.Set("diskEncryption", []interface{}{diskEnc})
}

// setPurchasingSection maps the 'purchasing' section of the computer inventory response to the Terraform schema.
func setPurchasingSection(d *schema.ResourceData, purchasing jamfpro.ComputerInventoryDataSubsetPurchasing) error {
	// Initialize a map to hold the 'purchasing' section attributes.
	purch := make(map[string]interface{})

	// Map each attribute of the 'purchasing' section from the API response to the corresponding Terraform schema attribute.
	purch["leased"] = purchasing.Leased
	purch["purchased"] = purchasing.Purchased
	purch["poNumber"] = purchasing.PoNumber
	purch["poDate"] = purchasing.PoDate
	purch["vendor"] = purchasing.Vendor
	purch["warrantyDate"] = purchasing.WarrantyDate
	purch["appleCareId"] = purchasing.AppleCareId
	purch["leaseDate"] = purchasing.LeaseDate
	purch["purchasePrice"] = purchasing.PurchasePrice
	purch["lifeExpectancy"] = purchasing.LifeExpectancy
	purch["purchasingAccount"] = purchasing.PurchasingAccount
	purch["purchasingContact"] = purchasing.PurchasingContact

	// Map 'extensionAttributes' as a list of maps.
	extAttrs := make([]map[string]interface{}, len(purchasing.ExtensionAttributes))
	for i, attr := range purchasing.ExtensionAttributes {
		attrMap := make(map[string]interface{})
		attrMap["definitionId"] = attr.DefinitionId
		attrMap["name"] = attr.Name
		attrMap["description"] = attr.Description
		attrMap["enabled"] = attr.Enabled
		attrMap["multiValue"] = attr.MultiValue
		attrMap["values"] = attr.Values
		attrMap["dataType"] = attr.DataType
		attrMap["options"] = attr.Options
		attrMap["inputType"] = attr.InputType

		extAttrs[i] = attrMap
	}
	purch["extensionAttributes"] = extAttrs

	// Set the 'purchasing' section in the Terraform resource data.
	return d.Set("purchasing", []interface{}{purch})
}

// setApplicationsSection maps the 'applications' section of the computer inventory response to the Terraform schema.
func setApplicationsSection(d *schema.ResourceData, applications []jamfpro.ComputerInventoryDataSubsetApplication) error {
	// Create a slice to hold the application maps.
	apps := make([]interface{}, len(applications))

	for i, app := range applications {
		// Initialize a map for each application.
		appMap := make(map[string]interface{})

		// Map each attribute of the application from the API response to the corresponding Terraform schema attribute.
		appMap["name"] = app.Name
		appMap["path"] = app.Path
		appMap["version"] = app.Version
		appMap["macAppStore"] = app.MacAppStore
		appMap["sizeMegabytes"] = app.SizeMegabytes
		appMap["bundleId"] = app.BundleId
		appMap["updateAvailable"] = app.UpdateAvailable
		appMap["externalVersionId"] = app.ExternalVersionId

		// Add the application map to the slice.
		apps[i] = appMap
	}

	// Set the 'applications' section in the Terraform resource data.
	return d.Set("applications", apps)
}

// setStorageSection maps the 'storage' section of the computer inventory response to the Terraform schema.
func setStorageSection(d *schema.ResourceData, storage jamfpro.ComputerInventoryDataSubsetStorage) error {
	storageMap := make(map[string]interface{})

	storageMap["bootDriveAvailableSpaceMegabytes"] = storage.BootDriveAvailableSpaceMegabytes

	// Mapping 'disks' array
	disks := make([]interface{}, len(storage.Disks))
	for i, disk := range storage.Disks {
		diskMap := make(map[string]interface{})
		diskMap["id"] = disk.ID
		diskMap["device"] = disk.Device
		diskMap["model"] = disk.Model
		diskMap["revision"] = disk.Revision
		diskMap["serialNumber"] = disk.SerialNumber
		diskMap["sizeMegabytes"] = disk.SizeMegabytes
		diskMap["smartStatus"] = disk.SmartStatus
		diskMap["type"] = disk.Type

		// Map 'partitions' if present
		partitions := make([]interface{}, len(disk.Partitions))
		for j, partition := range disk.Partitions {
			partitionMap := make(map[string]interface{})
			partitionMap["name"] = partition.Name
			partitionMap["sizeMegabytes"] = partition.SizeMegabytes
			partitionMap["availableMegabytes"] = partition.AvailableMegabytes
			partitionMap["partitionType"] = partition.PartitionType
			partitionMap["percentUsed"] = partition.PercentUsed
			partitionMap["fileVault2State"] = partition.FileVault2State
			partitionMap["fileVault2ProgressPercent"] = partition.FileVault2ProgressPercent
			partitionMap["lvmManaged"] = partition.LvmManaged
			partitions[j] = partitionMap
		}
		diskMap["partitions"] = partitions

		disks[i] = diskMap
	}
	storageMap["disks"] = disks

	// Set the 'storage' section in the Terraform resource data.
	return d.Set("storage", []interface{}{storageMap})
}

// setUserAndLocationSection maps the 'userAndLocation' section of the computer inventory response to the Terraform schema.
func setUserAndLocationSection(d *schema.ResourceData, userAndLocation jamfpro.ComputerInventoryDataSubsetUserAndLocation) error {
	userLocationMap := make(map[string]interface{})

	// Map each attribute from the 'userAndLocation' object to the corresponding schema attribute
	userLocationMap["username"] = userAndLocation.Username
	userLocationMap["realname"] = userAndLocation.Realname
	userLocationMap["email"] = userAndLocation.Email
	userLocationMap["position"] = userAndLocation.Position
	userLocationMap["phone"] = userAndLocation.Phone
	userLocationMap["departmentId"] = userAndLocation.DepartmentId
	userLocationMap["buildingId"] = userAndLocation.BuildingId
	userLocationMap["room"] = userAndLocation.Room

	// Map extension attributes if present
	if len(userAndLocation.ExtensionAttributes) > 0 {
		extAttrs := make([]map[string]interface{}, len(userAndLocation.ExtensionAttributes))
		for i, attr := range userAndLocation.ExtensionAttributes {
			attrMap := make(map[string]interface{})
			attrMap["definitionId"] = attr.DefinitionId
			attrMap["name"] = attr.Name
			attrMap["description"] = attr.Description
			attrMap["enabled"] = attr.Enabled
			attrMap["multiValue"] = attr.MultiValue
			attrMap["values"] = attr.Values
			attrMap["dataType"] = attr.DataType
			attrMap["options"] = attr.Options
			attrMap["inputType"] = attr.InputType

			extAttrs[i] = attrMap
		}
		userLocationMap["extensionAttributes"] = extAttrs
	}

	// Set the 'userAndLocation' section in the Terraform resource data
	return d.Set("userAndLocation", []interface{}{userLocationMap})
}

// setHardwareSection maps the 'hardware' section of the computer inventory response to the Terraform schema.
func setHardwareSection(d *schema.ResourceData, hardware jamfpro.ComputerInventoryDataSubsetHardware) error {
	hardwareMap := make(map[string]interface{})

	// Map each attribute from the 'hardware' object to the corresponding schema attribute
	hardwareMap["make"] = hardware.Make
	hardwareMap["model"] = hardware.Model
	hardwareMap["modelIdentifier"] = hardware.ModelIdentifier
	hardwareMap["serialNumber"] = hardware.SerialNumber
	hardwareMap["processorSpeedMhz"] = hardware.ProcessorSpeedMhz
	hardwareMap["processorCount"] = hardware.ProcessorCount
	hardwareMap["coreCount"] = hardware.CoreCount
	hardwareMap["processorType"] = hardware.ProcessorType
	hardwareMap["processorArchitecture"] = hardware.ProcessorArchitecture
	hardwareMap["busSpeedMhz"] = hardware.BusSpeedMhz
	hardwareMap["cacheSizeKilobytes"] = hardware.CacheSizeKilobytes
	hardwareMap["networkAdapterType"] = hardware.NetworkAdapterType
	hardwareMap["macAddress"] = hardware.MacAddress
	hardwareMap["altNetworkAdapterType"] = hardware.AltNetworkAdapterType
	hardwareMap["altMacAddress"] = hardware.AltMacAddress
	hardwareMap["totalRamMegabytes"] = hardware.TotalRamMegabytes
	hardwareMap["openRamSlots"] = hardware.OpenRamSlots
	hardwareMap["batteryCapacityPercent"] = hardware.BatteryCapacityPercent
	hardwareMap["smcVersion"] = hardware.SmcVersion
	hardwareMap["nicSpeed"] = hardware.NicSpeed
	hardwareMap["opticalDrive"] = hardware.OpticalDrive
	hardwareMap["bootRom"] = hardware.BootRom
	hardwareMap["bleCapable"] = hardware.BleCapable
	hardwareMap["supportsIosAppInstalls"] = hardware.SupportsIosAppInstalls
	hardwareMap["appleSilicon"] = hardware.AppleSilicon

	// Map extension attributes if present
	if len(hardware.ExtensionAttributes) > 0 {
		extAttrs := make([]map[string]interface{}, len(hardware.ExtensionAttributes))
		for i, attr := range hardware.ExtensionAttributes {
			attrMap := make(map[string]interface{})
			attrMap["definitionId"] = attr.DefinitionId
			attrMap["name"] = attr.Name
			attrMap["description"] = attr.Description
			attrMap["enabled"] = attr.Enabled
			attrMap["multiValue"] = attr.MultiValue
			attrMap["values"] = attr.Values
			attrMap["dataType"] = attr.DataType
			attrMap["options"] = attr.Options
			attrMap["inputType"] = attr.InputType

			extAttrs[i] = attrMap
		}
		hardwareMap["extensionAttributes"] = extAttrs
	}

	// Set the 'hardware' section in the Terraform resource data
	return d.Set("hardware", []interface{}{hardwareMap})
}

// setHardwareSection maps the 'hardware' section of the computer inventory response to the Terraform schema.
func setLocalUserAccountsSection(d *schema.ResourceData, localUserAccounts []jamfpro.ComputerInventoryDataSubsetLocalUserAccount) error {
	accounts := make([]interface{}, len(localUserAccounts))
	for i, account := range localUserAccounts {
		acc := make(map[string]interface{})
		acc["uid"] = account.UID
		acc["userGuid"] = account.UserGuid
		acc["username"] = account.Username
		acc["fullName"] = account.FullName
		acc["admin"] = account.Admin
		acc["homeDirectory"] = account.HomeDirectory
		acc["homeDirectorySizeMb"] = account.HomeDirectorySizeMb
		acc["fileVault2Enabled"] = account.FileVault2Enabled
		acc["userAccountType"] = account.UserAccountType
		acc["passwordMinLength"] = account.PasswordMinLength
		acc["passwordMaxAge"] = account.PasswordMaxAge
		acc["passwordMinComplexCharacters"] = account.PasswordMinComplexCharacters
		acc["passwordHistoryDepth"] = account.PasswordHistoryDepth
		acc["passwordRequireAlphanumeric"] = account.PasswordRequireAlphanumeric
		acc["computerAzureActiveDirectoryId"] = account.ComputerAzureActiveDirectoryId
		acc["userAzureActiveDirectoryId"] = account.UserAzureActiveDirectoryId
		acc["azureActiveDirectoryId"] = account.AzureActiveDirectoryId
		accounts[i] = acc
	}
	return d.Set("localUserAccounts", accounts)
}

// setCertificatesSection maps the 'certificate' section of the computer inventory response to the Terraform schema.
func setCertificatesSection(d *schema.ResourceData, certificates []jamfpro.ComputerInventoryDataSubsetCertificate) error {
	certs := make([]interface{}, len(certificates))
	for i, cert := range certificates {
		certMap := make(map[string]interface{})
		certMap["commonName"] = cert.CommonName
		certMap["identity"] = cert.Identity
		certMap["expirationDate"] = cert.ExpirationDate
		certMap["username"] = cert.Username
		certMap["lifecycleStatus"] = cert.LifecycleStatus
		certMap["certificateStatus"] = cert.CertificateStatus
		certMap["subjectName"] = cert.SubjectName
		certMap["serialNumber"] = cert.SerialNumber
		certMap["sha1Fingerprint"] = cert.Sha1Fingerprint
		certMap["issuedDate"] = cert.IssuedDate
		certs[i] = certMap
	}
	return d.Set("certificates", certs)
}

// setAttachmentsSection maps the 'attachments' section of the computer inventory response to the Terraform schema.
func setAttachmentsSection(d *schema.ResourceData, attachments []jamfpro.ComputerInventoryDataSubsetAttachment) error {
	atts := make([]interface{}, len(attachments))
	for i, att := range attachments {
		attMap := make(map[string]interface{})
		attMap["id"] = att.ID
		attMap["name"] = att.Name
		attMap["fileType"] = att.FileType
		attMap["sizeBytes"] = att.SizeBytes
		atts[i] = attMap
	}
	return d.Set("attachments", atts)
}

// setPluginsSection maps the 'plugins' section of the computer inventory response to the Terraform schema.
func setPluginsSection(d *schema.ResourceData, plugins []jamfpro.ComputerInventoryDataSubsetPlugin) error {
	pluginList := make([]interface{}, len(plugins))
	for i, plugin := range plugins {
		pluginMap := make(map[string]interface{})
		pluginMap["name"] = plugin.Name
		pluginMap["version"] = plugin.Version
		pluginMap["path"] = plugin.Path
		pluginList[i] = pluginMap
	}
	return d.Set("plugins", pluginList)
}

// setPackageReceiptsSection maps the 'package receipts' section of the computer inventory response to the Terraform schema.
func setPackageReceiptsSection(d *schema.ResourceData, packageReceipts jamfpro.ComputerInventoryDataSubsetPackageReceipts) error {
	prMap := make(map[string]interface{})
	prMap["installedByJamfPro"] = packageReceipts.InstalledByJamfPro
	prMap["installedByInstallerSwu"] = packageReceipts.InstalledByInstallerSwu
	prMap["cached"] = packageReceipts.Cached
	return d.Set("packageReceipts", []interface{}{prMap})
}

// setFontsSection maps the 'fonts' section of the computer inventory response to the Terraform schema.
func setFontsSection(d *schema.ResourceData, fonts []jamfpro.ComputerInventoryDataSubsetFont) error {
	fontsList := make([]interface{}, len(fonts))
	for i, font := range fonts {
		fontMap := make(map[string]interface{})
		fontMap["name"] = font.Name
		fontMap["version"] = font.Version
		fontMap["path"] = font.Path
		fontsList[i] = fontMap
	}
	return d.Set("fonts", fontsList)
}

// setSecuritySection maps the 'security' section of the computer inventory response to the Terraform schema.
func setSecuritySection(d *schema.ResourceData, security jamfpro.ComputerInventoryDataSubsetSecurity) error {
	securityMap := make(map[string]interface{})
	securityMap["sipStatus"] = security.SipStatus
	securityMap["gatekeeperStatus"] = security.GatekeeperStatus
	securityMap["xprotectVersion"] = security.XprotectVersion
	securityMap["autoLoginDisabled"] = security.AutoLoginDisabled
	securityMap["remoteDesktopEnabled"] = security.RemoteDesktopEnabled
	securityMap["activationLockEnabled"] = security.ActivationLockEnabled
	securityMap["recoveryLockEnabled"] = security.RecoveryLockEnabled
	securityMap["firewallEnabled"] = security.FirewallEnabled
	securityMap["secureBootLevel"] = security.SecureBootLevel
	securityMap["externalBootLevel"] = security.ExternalBootLevel
	securityMap["bootstrapTokenAllowed"] = security.BootstrapTokenAllowed
	return d.Set("security", []interface{}{securityMap})
}

// setOperatingSystemSection maps the 'Operating System' section of the computer inventory response to the Terraform schema.
func setOperatingSystemSection(d *schema.ResourceData, operatingSystem jamfpro.ComputerInventoryDataSubsetOperatingSystem) error {
	osMap := make(map[string]interface{})
	osMap["name"] = operatingSystem.Name
	osMap["version"] = operatingSystem.Version
	osMap["build"] = operatingSystem.Build
	osMap["supplementalBuildVersion"] = operatingSystem.SupplementalBuildVersion
	osMap["rapidSecurityResponse"] = operatingSystem.RapidSecurityResponse
	osMap["activeDirectoryStatus"] = operatingSystem.ActiveDirectoryStatus
	osMap["fileVault2Status"] = operatingSystem.FileVault2Status
	osMap["softwareUpdateDeviceId"] = operatingSystem.SoftwareUpdateDeviceId
	// Map extension attributes if present
	extAttrs := make([]map[string]interface{}, len(operatingSystem.ExtensionAttributes))
	for i, attr := range operatingSystem.ExtensionAttributes {
		attrMap := make(map[string]interface{})
		attrMap["definitionId"] = attr.DefinitionId
		attrMap["name"] = attr.Name
		attrMap["description"] = attr.Description
		attrMap["enabled"] = attr.Enabled
		attrMap["multiValue"] = attr.MultiValue
		attrMap["values"] = attr.Values
		attrMap["dataType"] = attr.DataType
		attrMap["options"] = attr.Options
		attrMap["inputType"] = attr.InputType

		extAttrs[i] = attrMap
	}
	osMap["extensionAttributes"] = extAttrs
	return d.Set("operatingSystem", []interface{}{osMap})
}

// setLicensedSoftwareSection maps the 'Licensed Software' section of the computer inventory response to the Terraform schema.
func setLicensedSoftwareSection(d *schema.ResourceData, licensedSoftware []jamfpro.ComputerInventoryDataSubsetLicensedSoftware) error {
	softwareList := make([]interface{}, len(licensedSoftware))
	for i, software := range licensedSoftware {
		softwareMap := make(map[string]interface{})
		softwareMap["id"] = software.ID
		softwareMap["name"] = software.Name
		softwareList[i] = softwareMap
	}
	return d.Set("licensedSoftware", softwareList)
}

// setIBeaconsSection maps the 'IBeacons' section of the computer inventory response to the Terraform schema.
func setIBeaconsSection(d *schema.ResourceData, ibeacons []jamfpro.ComputerInventoryDataSubsetIbeacon) error {
	ibeaconList := make([]interface{}, len(ibeacons))
	for i, ibeacon := range ibeacons {
		ibeaconMap := make(map[string]interface{})
		ibeaconMap["name"] = ibeacon.Name
		ibeaconList[i] = ibeaconMap
	}
	return d.Set("ibeacons", ibeaconList)
}

// setSoftwareUpdatesSection maps the 'Software Updates' section of the computer inventory response to the Terraform schema.
func setSoftwareUpdatesSection(d *schema.ResourceData, softwareUpdates []jamfpro.ComputerInventoryDataSubsetSoftwareUpdate) error {
	updateList := make([]interface{}, len(softwareUpdates))
	for i, update := range softwareUpdates {
		updateMap := make(map[string]interface{})
		updateMap["name"] = update.Name
		updateMap["version"] = update.Version
		updateMap["packageName"] = update.PackageName
		updateList[i] = updateMap
	}
	return d.Set("softwareUpdates", updateList)
}

// setExtensionAttributesSection maps the 'Extension Attributes' section of the computer inventory response to the Terraform schema.
func setExtensionAttributesSection(d *schema.ResourceData, extensionAttributes []jamfpro.ComputerInventoryDataSubsetExtensionAttribute) error {
	attrList := make([]interface{}, len(extensionAttributes))
	for i, attr := range extensionAttributes {
		attrMap := make(map[string]interface{})
		attrMap["definitionId"] = attr.DefinitionId
		attrMap["name"] = attr.Name
		attrMap["description"] = attr.Description
		attrMap["enabled"] = attr.Enabled
		attrMap["multiValue"] = attr.MultiValue
		attrMap["values"] = attr.Values
		attrMap["dataType"] = attr.DataType
		attrMap["options"] = attr.Options
		attrMap["inputType"] = attr.InputType
		attrList[i] = attrMap
	}
	return d.Set("extensionAttributes", attrList)
}

// setGroupMembershipsSection maps the 'groupMemberships' section of the computer inventory response to the Terraform schema.
func setGroupMembershipsSection(d *schema.ResourceData, groupMemberships []jamfpro.ComputerInventoryDataSubsetGroupMembership) error {
	memberships := make([]interface{}, len(groupMemberships))
	for i, group := range groupMemberships {
		groupMap := make(map[string]interface{})
		groupMap["groupId"] = group.GroupId
		groupMap["groupName"] = group.GroupName
		groupMap["smartGroup"] = group.SmartGroup

		memberships[i] = groupMap
	}
	return d.Set("groupMemberships", memberships)
}
