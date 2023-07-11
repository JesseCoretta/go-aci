// Code generated from ACIParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ACIParser
import "github.com/antlr4-go/antlr/v4"

// ACIParserListener is a complete listener for a parse tree produced by ACIParser.
type ACIParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterAci is called when entering the aci production.
	EnterAci(c *AciContext)

	// EnterPermission_bind_rules is called when entering the permission_bind_rules production.
	EnterPermission_bind_rules(c *Permission_bind_rulesContext)

	// EnterPermission_and_bind_rule_pair is called when entering the permission_and_bind_rule_pair production.
	EnterPermission_and_bind_rule_pair(c *Permission_and_bind_rule_pairContext)

	// EnterPermission_expression is called when entering the permission_expression production.
	EnterPermission_expression(c *Permission_expressionContext)

	// EnterAllow_access is called when entering the allow_access production.
	EnterAllow_access(c *Allow_accessContext)

	// EnterDeny_access is called when entering the deny_access production.
	EnterDeny_access(c *Deny_accessContext)

	// EnterSearch_privilege is called when entering the search_privilege production.
	EnterSearch_privilege(c *Search_privilegeContext)

	// EnterRead_privilege is called when entering the read_privilege production.
	EnterRead_privilege(c *Read_privilegeContext)

	// EnterCompare_privilege is called when entering the compare_privilege production.
	EnterCompare_privilege(c *Compare_privilegeContext)

	// EnterAdd_privilege is called when entering the add_privilege production.
	EnterAdd_privilege(c *Add_privilegeContext)

	// EnterDelete_privilege is called when entering the delete_privilege production.
	EnterDelete_privilege(c *Delete_privilegeContext)

	// EnterSelfwrite_privilege is called when entering the selfwrite_privilege production.
	EnterSelfwrite_privilege(c *Selfwrite_privilegeContext)

	// EnterProxy_privilege is called when entering the proxy_privilege production.
	EnterProxy_privilege(c *Proxy_privilegeContext)

	// EnterImport_privilege is called when entering the import_privilege production.
	EnterImport_privilege(c *Import_privilegeContext)

	// EnterExport_privilege is called when entering the export_privilege production.
	EnterExport_privilege(c *Export_privilegeContext)

	// EnterAll_privileges is called when entering the all_privileges production.
	EnterAll_privileges(c *All_privilegesContext)

	// EnterTarget_rules is called when entering the target_rules production.
	EnterTarget_rules(c *Target_rulesContext)

	// EnterTargetRule is called when entering the targetRule production.
	EnterTargetRule(c *TargetRuleContext)

	// EnterTarget_dn_rule is called when entering the target_dn_rule production.
	EnterTarget_dn_rule(c *Target_dn_ruleContext)

	// EnterTarget_to_rule is called when entering the target_to_rule production.
	EnterTarget_to_rule(c *Target_to_ruleContext)

	// EnterTarget_from_rule is called when entering the target_from_rule production.
	EnterTarget_from_rule(c *Target_from_ruleContext)

	// EnterTargetfilter_rule is called when entering the targetfilter_rule production.
	EnterTargetfilter_rule(c *Targetfilter_ruleContext)

	// EnterTargattrfilters_rule is called when entering the targattrfilters_rule production.
	EnterTargattrfilters_rule(c *Targattrfilters_ruleContext)

	// EnterTargetscope_rule is called when entering the targetscope_rule production.
	EnterTargetscope_rule(c *Targetscope_ruleContext)

	// EnterTargetattr_rule is called when entering the targetattr_rule production.
	EnterTargetattr_rule(c *Targetattr_ruleContext)

	// EnterTargetcontrol_rule is called when entering the targetcontrol_rule production.
	EnterTargetcontrol_rule(c *Targetcontrol_ruleContext)

	// EnterTargetextop_rule is called when entering the targetextop_rule production.
	EnterTargetextop_rule(c *Targetextop_ruleContext)

	// EnterBase_object_targetscope is called when entering the base_object_targetscope production.
	EnterBase_object_targetscope(c *Base_object_targetscopeContext)

	// EnterOne_level_targetscope is called when entering the one_level_targetscope production.
	EnterOne_level_targetscope(c *One_level_targetscopeContext)

	// EnterSub_tree_targetscope is called when entering the sub_tree_targetscope production.
	EnterSub_tree_targetscope(c *Sub_tree_targetscopeContext)

	// EnterSubordinate_targetscope is called when entering the subordinate_targetscope production.
	EnterSubordinate_targetscope(c *Subordinate_targetscopeContext)

	// EnterQuoted_object_identifier_list is called when entering the quoted_object_identifier_list production.
	EnterQuoted_object_identifier_list(c *Quoted_object_identifier_listContext)

	// EnterList_of_quoted_object_identifiers is called when entering the list_of_quoted_object_identifiers production.
	EnterList_of_quoted_object_identifiers(c *List_of_quoted_object_identifiersContext)

	// EnterQuoted_targeted_attributes_list is called when entering the quoted_targeted_attributes_list production.
	EnterQuoted_targeted_attributes_list(c *Quoted_targeted_attributes_listContext)

	// EnterList_of_quoted_attributes is called when entering the list_of_quoted_attributes production.
	EnterList_of_quoted_attributes(c *List_of_quoted_attributesContext)

	// EnterAll_attributes is called when entering the all_attributes production.
	EnterAll_attributes(c *All_attributesContext)

	// EnterObject_identifier is called when entering the object_identifier production.
	EnterObject_identifier(c *Object_identifierContext)

	// EnterNumber_form is called when entering the number_form production.
	EnterNumber_form(c *Number_formContext)

	// EnterAttribute_filters_sets is called when entering the attribute_filters_sets production.
	EnterAttribute_filters_sets(c *Attribute_filters_setsContext)

	// EnterAttribute_filters_set is called when entering the attribute_filters_set production.
	EnterAttribute_filters_set(c *Attribute_filters_setContext)

	// EnterAttribute_filter_single is called when entering the attribute_filter_single production.
	EnterAttribute_filter_single(c *Attribute_filter_singleContext)

	// EnterAttribute_filters is called when entering the attribute_filters production.
	EnterAttribute_filters(c *Attribute_filtersContext)

	// EnterAttribute_filter_set is called when entering the attribute_filter_set production.
	EnterAttribute_filter_set(c *Attribute_filter_setContext)

	// EnterAdd_filter_operation is called when entering the add_filter_operation production.
	EnterAdd_filter_operation(c *Add_filter_operationContext)

	// EnterDelete_filter_operation is called when entering the delete_filter_operation production.
	EnterDelete_filter_operation(c *Delete_filter_operationContext)

	// EnterAttribute_filter is called when entering the attribute_filter production.
	EnterAttribute_filter(c *Attribute_filterContext)

	// EnterBind_rule is called when entering the bind_rule production.
	EnterBind_rule(c *Bind_ruleContext)

	// EnterParenthetical_bind_rule is called when entering the parenthetical_bind_rule production.
	EnterParenthetical_bind_rule(c *Parenthetical_bind_ruleContext)

	// EnterParenthetical_bind_rule_req_bool_op is called when entering the parenthetical_bind_rule_req_bool_op production.
	EnterParenthetical_bind_rule_req_bool_op(c *Parenthetical_bind_rule_req_bool_opContext)

	// EnterNegated_bind_rule_expression is called when entering the negated_bind_rule_expression production.
	EnterNegated_bind_rule_expression(c *Negated_bind_rule_expressionContext)

	// EnterParenthetical_bind_rule_expression is called when entering the parenthetical_bind_rule_expression production.
	EnterParenthetical_bind_rule_expression(c *Parenthetical_bind_rule_expressionContext)

	// EnterBind_rule_expression_recursion is called when entering the bind_rule_expression_recursion production.
	EnterBind_rule_expression_recursion(c *Bind_rule_expression_recursionContext)

	// EnterBind_rule_expression is called when entering the bind_rule_expression production.
	EnterBind_rule_expression(c *Bind_rule_expressionContext)

	// EnterUserdn_expression is called when entering the userdn_expression production.
	EnterUserdn_expression(c *Userdn_expressionContext)

	// EnterUserattr_expression is called when entering the userattr_expression production.
	EnterUserattr_expression(c *Userattr_expressionContext)

	// EnterGroupdn_expression is called when entering the groupdn_expression production.
	EnterGroupdn_expression(c *Groupdn_expressionContext)

	// EnterGroupattr_expression is called when entering the groupattr_expression production.
	EnterGroupattr_expression(c *Groupattr_expressionContext)

	// EnterRoledn_expression is called when entering the roledn_expression production.
	EnterRoledn_expression(c *Roledn_expressionContext)

	// EnterDns_expression is called when entering the dns_expression production.
	EnterDns_expression(c *Dns_expressionContext)

	// EnterIp_expression is called when entering the ip_expression production.
	EnterIp_expression(c *Ip_expressionContext)

	// EnterTimeofday_expression is called when entering the timeofday_expression production.
	EnterTimeofday_expression(c *Timeofday_expressionContext)

	// EnterDayofweek_expression is called when entering the dayofweek_expression production.
	EnterDayofweek_expression(c *Dayofweek_expressionContext)

	// EnterSsf_expression is called when entering the ssf_expression production.
	EnterSsf_expression(c *Ssf_expressionContext)

	// EnterAuthmethod_expression is called when entering the authmethod_expression production.
	EnterAuthmethod_expression(c *Authmethod_expressionContext)

	// EnterParenthetical_dayofweek_bind_rule is called when entering the parenthetical_dayofweek_bind_rule production.
	EnterParenthetical_dayofweek_bind_rule(c *Parenthetical_dayofweek_bind_ruleContext)

	// EnterDayofweek_bind_rule is called when entering the dayofweek_bind_rule production.
	EnterDayofweek_bind_rule(c *Dayofweek_bind_ruleContext)

	// EnterSun is called when entering the Sun production.
	EnterSun(c *SunContext)

	// EnterMon is called when entering the Mon production.
	EnterMon(c *MonContext)

	// EnterTues is called when entering the Tues production.
	EnterTues(c *TuesContext)

	// EnterWed is called when entering the Wed production.
	EnterWed(c *WedContext)

	// EnterThur is called when entering the Thur production.
	EnterThur(c *ThurContext)

	// EnterFri is called when entering the Fri production.
	EnterFri(c *FriContext)

	// EnterSat is called when entering the Sat production.
	EnterSat(c *SatContext)

	// EnterParentheticalAuthenticationMethod is called when entering the parentheticalAuthenticationMethod production.
	EnterParentheticalAuthenticationMethod(c *ParentheticalAuthenticationMethodContext)

	// EnterAuthentication_method is called when entering the authentication_method production.
	EnterAuthentication_method(c *Authentication_methodContext)

	// EnterNone is called when entering the none production.
	EnterNone(c *NoneContext)

	// EnterSimple is called when entering the simple production.
	EnterSimple(c *SimpleContext)

	// EnterSsl is called when entering the ssl production.
	EnterSsl(c *SslContext)

	// EnterSasl is called when entering the sasl production.
	EnterSasl(c *SaslContext)

	// EnterParenthetical_bind_userdn is called when entering the parenthetical_bind_userdn production.
	EnterParenthetical_bind_userdn(c *Parenthetical_bind_userdnContext)

	// EnterBind_userdn is called when entering the bind_userdn production.
	EnterBind_userdn(c *Bind_userdnContext)

	// EnterParenthetical_bind_roledn is called when entering the parenthetical_bind_roledn production.
	EnterParenthetical_bind_roledn(c *Parenthetical_bind_rolednContext)

	// EnterBind_roledn is called when entering the bind_roledn production.
	EnterBind_roledn(c *Bind_rolednContext)

	// EnterParenthetical_bind_groupdn is called when entering the parenthetical_bind_groupdn production.
	EnterParenthetical_bind_groupdn(c *Parenthetical_bind_groupdnContext)

	// EnterBind_groupdn is called when entering the bind_groupdn production.
	EnterBind_groupdn(c *Bind_groupdnContext)

	// EnterParenthetical_bind_userattr is called when entering the parenthetical_bind_userattr production.
	EnterParenthetical_bind_userattr(c *Parenthetical_bind_userattrContext)

	// EnterBind_userattr is called when entering the bind_userattr production.
	EnterBind_userattr(c *Bind_userattrContext)

	// EnterParenthetical_bind_groupattr is called when entering the parenthetical_bind_groupattr production.
	EnterParenthetical_bind_groupattr(c *Parenthetical_bind_groupattrContext)

	// EnterBind_groupattr is called when entering the bind_groupattr production.
	EnterBind_groupattr(c *Bind_groupattrContext)

	// EnterParenthetical_ssf is called when entering the parenthetical_ssf production.
	EnterParenthetical_ssf(c *Parenthetical_ssfContext)

	// EnterBind_ssf is called when entering the bind_ssf production.
	EnterBind_ssf(c *Bind_ssfContext)

	// EnterParenthetical_bind_timeofday is called when entering the parenthetical_bind_timeofday production.
	EnterParenthetical_bind_timeofday(c *Parenthetical_bind_timeofdayContext)

	// EnterBind_timeofday is called when entering the bind_timeofday production.
	EnterBind_timeofday(c *Bind_timeofdayContext)

	// EnterParenthetical_bind_ip is called when entering the parenthetical_bind_ip production.
	EnterParenthetical_bind_ip(c *Parenthetical_bind_ipContext)

	// EnterBind_ip is called when entering the bind_ip production.
	EnterBind_ip(c *Bind_ipContext)

	// EnterParenthetical_bind_dns is called when entering the parenthetical_bind_dns production.
	EnterParenthetical_bind_dns(c *Parenthetical_bind_dnsContext)

	// EnterDns_bind_rule is called when entering the dns_bind_rule production.
	EnterDns_bind_rule(c *Dns_bind_ruleContext)

	// EnterIps is called when entering the ips production.
	EnterIps(c *IpsContext)

	// EnterIpv4_address is called when entering the ipv4_address production.
	EnterIpv4_address(c *Ipv4_addressContext)

	// EnterIpv6_address is called when entering the ipv6_address production.
	EnterIpv6_address(c *Ipv6_addressContext)

	// EnterIpv4 is called when entering the ipv4 production.
	EnterIpv4(c *Ipv4Context)

	// EnterIpv6 is called when entering the ipv6 production.
	EnterIpv6(c *Ipv6Context)

	// EnterFqdn is called when entering the fqdn production.
	EnterFqdn(c *FqdnContext)

	// EnterFullyQualifiedLDAPURI is called when entering the fullyQualifiedLDAPURI production.
	EnterFullyQualifiedLDAPURI(c *FullyQualifiedLDAPURIContext)

	// EnterUriSearchFilter is called when entering the uriSearchFilter production.
	EnterUriSearchFilter(c *UriSearchFilterContext)

	// EnterUriSearchScopes is called when entering the uriSearchScopes production.
	EnterUriSearchScopes(c *UriSearchScopesContext)

	// EnterUriAttributeList is called when entering the uriAttributeList production.
	EnterUriAttributeList(c *UriAttributeListContext)

	// EnterQuoted_distinguished_name_list is called when entering the quoted_distinguished_name_list production.
	EnterQuoted_distinguished_name_list(c *Quoted_distinguished_name_listContext)

	// EnterList_of_quoted_distinguished_names is called when entering the list_of_quoted_distinguished_names production.
	EnterList_of_quoted_distinguished_names(c *List_of_quoted_distinguished_namesContext)

	// EnterDn is called when entering the dn production.
	EnterDn(c *DnContext)

	// EnterAnonymous_dn_alias is called when entering the anonymous_dn_alias production.
	EnterAnonymous_dn_alias(c *Anonymous_dn_aliasContext)

	// EnterAny_user_dn_alias is called when entering the any_user_dn_alias production.
	EnterAny_user_dn_alias(c *Any_user_dn_aliasContext)

	// EnterSelf_dn_alias is called when entering the self_dn_alias production.
	EnterSelf_dn_alias(c *Self_dn_aliasContext)

	// EnterParent_dn_alias is called when entering the parent_dn_alias production.
	EnterParent_dn_alias(c *Parent_dn_aliasContext)

	// EnterRdn_macro is called when entering the rdn_macro production.
	EnterRdn_macro(c *Rdn_macroContext)

	// EnterParenthetical_filter_expression is called when entering the parenthetical_filter_expression production.
	EnterParenthetical_filter_expression(c *Parenthetical_filter_expressionContext)

	// EnterFilter_expressions is called when entering the filter_expressions production.
	EnterFilter_expressions(c *Filter_expressionsContext)

	// EnterParenthetical_filter_expression_opt_bool is called when entering the parenthetical_filter_expression_opt_bool production.
	EnterParenthetical_filter_expression_opt_bool(c *Parenthetical_filter_expression_opt_boolContext)

	// EnterNot_filter_expression is called when entering the not_filter_expression production.
	EnterNot_filter_expression(c *Not_filter_expressionContext)

	// EnterAva_expression is called when entering the ava_expression production.
	EnterAva_expression(c *Ava_expressionContext)

	// EnterAva_or_rdn is called when entering the ava_or_rdn production.
	EnterAva_or_rdn(c *Ava_or_rdnContext)

	// EnterInheritance_expression is called when entering the inheritance_expression production.
	EnterInheritance_expression(c *Inheritance_expressionContext)

	// EnterInheritance_levels is called when entering the inheritance_levels production.
	EnterInheritance_levels(c *Inheritance_levelsContext)

	// EnterAttr_bind_type_or_value is called when entering the attr_bind_type_or_value production.
	EnterAttr_bind_type_or_value(c *Attr_bind_type_or_valueContext)

	// EnterUSERDN is called when entering the USERDN production.
	EnterUSERDN(c *USERDNContext)

	// EnterGROUPDN is called when entering the GROUPDN production.
	EnterGROUPDN(c *GROUPDNContext)

	// EnterROLEDN is called when entering the ROLEDN production.
	EnterROLEDN(c *ROLEDNContext)

	// EnterSELFDN is called when entering the SELFDN production.
	EnterSELFDN(c *SELFDNContext)

	// EnterLDAPURL is called when entering the LDAPURL production.
	EnterLDAPURL(c *LDAPURLContext)

	// EnterKey_or_value is called when entering the key_or_value production.
	EnterKey_or_value(c *Key_or_valueContext)

	// EnterEqual_to is called when entering the equal_to production.
	EnterEqual_to(c *Equal_toContext)

	// EnterNot_equal_to is called when entering the not_equal_to production.
	EnterNot_equal_to(c *Not_equal_toContext)

	// EnterGreater_than_or_equal is called when entering the greater_than_or_equal production.
	EnterGreater_than_or_equal(c *Greater_than_or_equalContext)

	// EnterLess_than_or_equal is called when entering the less_than_or_equal production.
	EnterLess_than_or_equal(c *Less_than_or_equalContext)

	// EnterApprox is called when entering the approx production.
	EnterApprox(c *ApproxContext)

	// EnterExtensible_rule is called when entering the extensible_rule production.
	EnterExtensible_rule(c *Extensible_ruleContext)

	// EnterExtensible_rule_with_dn is called when entering the extensible_rule_with_dn production.
	EnterExtensible_rule_with_dn(c *Extensible_rule_with_dnContext)

	// EnterExtensible_rule_with_attrs is called when entering the extensible_rule_with_attrs production.
	EnterExtensible_rule_with_attrs(c *Extensible_rule_with_attrsContext)

	// EnterExtensible_rule_with_dn_oid is called when entering the extensible_rule_with_dn_oid production.
	EnterExtensible_rule_with_dn_oid(c *Extensible_rule_with_dn_oidContext)

	// EnterEqualTo is called when entering the equalTo production.
	EnterEqualTo(c *EqualToContext)

	// EnterNotEqualTo is called when entering the notEqualTo production.
	EnterNotEqualTo(c *NotEqualToContext)

	// EnterGreaterThan is called when entering the greaterThan production.
	EnterGreaterThan(c *GreaterThanContext)

	// EnterLessThan is called when entering the lessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterGreaterThanOrEqual is called when entering the greaterThanOrEqual production.
	EnterGreaterThanOrEqual(c *GreaterThanOrEqualContext)

	// EnterLessThanOrEqual is called when entering the lessThanOrEqual production.
	EnterLessThanOrEqual(c *LessThanOrEqualContext)

	// EnterApproximate is called when entering the approximate production.
	EnterApproximate(c *ApproximateContext)

	// EnterExtensibleRule is called when entering the extensibleRule production.
	EnterExtensibleRule(c *ExtensibleRuleContext)

	// EnterExtensibleRuleDNOID is called when entering the extensibleRuleDNOID production.
	EnterExtensibleRuleDNOID(c *ExtensibleRuleDNOIDContext)

	// EnterExtensibleRuleDN is called when entering the extensibleRuleDN production.
	EnterExtensibleRuleDN(c *ExtensibleRuleDNContext)

	// EnterExtensibleRuleAttr is called when entering the extensibleRuleAttr production.
	EnterExtensibleRuleAttr(c *ExtensibleRuleAttrContext)

	// EnterORDelimiter is called when entering the oRDelimiter production.
	EnterORDelimiter(c *ORDelimiterContext)

	// EnterANDDelimiter is called when entering the aNDDelimiter production.
	EnterANDDelimiter(c *ANDDelimiterContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitAci is called when exiting the aci production.
	ExitAci(c *AciContext)

	// ExitPermission_bind_rules is called when exiting the permission_bind_rules production.
	ExitPermission_bind_rules(c *Permission_bind_rulesContext)

	// ExitPermission_and_bind_rule_pair is called when exiting the permission_and_bind_rule_pair production.
	ExitPermission_and_bind_rule_pair(c *Permission_and_bind_rule_pairContext)

	// ExitPermission_expression is called when exiting the permission_expression production.
	ExitPermission_expression(c *Permission_expressionContext)

	// ExitAllow_access is called when exiting the allow_access production.
	ExitAllow_access(c *Allow_accessContext)

	// ExitDeny_access is called when exiting the deny_access production.
	ExitDeny_access(c *Deny_accessContext)

	// ExitSearch_privilege is called when exiting the search_privilege production.
	ExitSearch_privilege(c *Search_privilegeContext)

	// ExitRead_privilege is called when exiting the read_privilege production.
	ExitRead_privilege(c *Read_privilegeContext)

	// ExitCompare_privilege is called when exiting the compare_privilege production.
	ExitCompare_privilege(c *Compare_privilegeContext)

	// ExitAdd_privilege is called when exiting the add_privilege production.
	ExitAdd_privilege(c *Add_privilegeContext)

	// ExitDelete_privilege is called when exiting the delete_privilege production.
	ExitDelete_privilege(c *Delete_privilegeContext)

	// ExitSelfwrite_privilege is called when exiting the selfwrite_privilege production.
	ExitSelfwrite_privilege(c *Selfwrite_privilegeContext)

	// ExitProxy_privilege is called when exiting the proxy_privilege production.
	ExitProxy_privilege(c *Proxy_privilegeContext)

	// ExitImport_privilege is called when exiting the import_privilege production.
	ExitImport_privilege(c *Import_privilegeContext)

	// ExitExport_privilege is called when exiting the export_privilege production.
	ExitExport_privilege(c *Export_privilegeContext)

	// ExitAll_privileges is called when exiting the all_privileges production.
	ExitAll_privileges(c *All_privilegesContext)

	// ExitTarget_rules is called when exiting the target_rules production.
	ExitTarget_rules(c *Target_rulesContext)

	// ExitTargetRule is called when exiting the targetRule production.
	ExitTargetRule(c *TargetRuleContext)

	// ExitTarget_dn_rule is called when exiting the target_dn_rule production.
	ExitTarget_dn_rule(c *Target_dn_ruleContext)

	// ExitTarget_to_rule is called when exiting the target_to_rule production.
	ExitTarget_to_rule(c *Target_to_ruleContext)

	// ExitTarget_from_rule is called when exiting the target_from_rule production.
	ExitTarget_from_rule(c *Target_from_ruleContext)

	// ExitTargetfilter_rule is called when exiting the targetfilter_rule production.
	ExitTargetfilter_rule(c *Targetfilter_ruleContext)

	// ExitTargattrfilters_rule is called when exiting the targattrfilters_rule production.
	ExitTargattrfilters_rule(c *Targattrfilters_ruleContext)

	// ExitTargetscope_rule is called when exiting the targetscope_rule production.
	ExitTargetscope_rule(c *Targetscope_ruleContext)

	// ExitTargetattr_rule is called when exiting the targetattr_rule production.
	ExitTargetattr_rule(c *Targetattr_ruleContext)

	// ExitTargetcontrol_rule is called when exiting the targetcontrol_rule production.
	ExitTargetcontrol_rule(c *Targetcontrol_ruleContext)

	// ExitTargetextop_rule is called when exiting the targetextop_rule production.
	ExitTargetextop_rule(c *Targetextop_ruleContext)

	// ExitBase_object_targetscope is called when exiting the base_object_targetscope production.
	ExitBase_object_targetscope(c *Base_object_targetscopeContext)

	// ExitOne_level_targetscope is called when exiting the one_level_targetscope production.
	ExitOne_level_targetscope(c *One_level_targetscopeContext)

	// ExitSub_tree_targetscope is called when exiting the sub_tree_targetscope production.
	ExitSub_tree_targetscope(c *Sub_tree_targetscopeContext)

	// ExitSubordinate_targetscope is called when exiting the subordinate_targetscope production.
	ExitSubordinate_targetscope(c *Subordinate_targetscopeContext)

	// ExitQuoted_object_identifier_list is called when exiting the quoted_object_identifier_list production.
	ExitQuoted_object_identifier_list(c *Quoted_object_identifier_listContext)

	// ExitList_of_quoted_object_identifiers is called when exiting the list_of_quoted_object_identifiers production.
	ExitList_of_quoted_object_identifiers(c *List_of_quoted_object_identifiersContext)

	// ExitQuoted_targeted_attributes_list is called when exiting the quoted_targeted_attributes_list production.
	ExitQuoted_targeted_attributes_list(c *Quoted_targeted_attributes_listContext)

	// ExitList_of_quoted_attributes is called when exiting the list_of_quoted_attributes production.
	ExitList_of_quoted_attributes(c *List_of_quoted_attributesContext)

	// ExitAll_attributes is called when exiting the all_attributes production.
	ExitAll_attributes(c *All_attributesContext)

	// ExitObject_identifier is called when exiting the object_identifier production.
	ExitObject_identifier(c *Object_identifierContext)

	// ExitNumber_form is called when exiting the number_form production.
	ExitNumber_form(c *Number_formContext)

	// ExitAttribute_filters_sets is called when exiting the attribute_filters_sets production.
	ExitAttribute_filters_sets(c *Attribute_filters_setsContext)

	// ExitAttribute_filters_set is called when exiting the attribute_filters_set production.
	ExitAttribute_filters_set(c *Attribute_filters_setContext)

	// ExitAttribute_filter_single is called when exiting the attribute_filter_single production.
	ExitAttribute_filter_single(c *Attribute_filter_singleContext)

	// ExitAttribute_filters is called when exiting the attribute_filters production.
	ExitAttribute_filters(c *Attribute_filtersContext)

	// ExitAttribute_filter_set is called when exiting the attribute_filter_set production.
	ExitAttribute_filter_set(c *Attribute_filter_setContext)

	// ExitAdd_filter_operation is called when exiting the add_filter_operation production.
	ExitAdd_filter_operation(c *Add_filter_operationContext)

	// ExitDelete_filter_operation is called when exiting the delete_filter_operation production.
	ExitDelete_filter_operation(c *Delete_filter_operationContext)

	// ExitAttribute_filter is called when exiting the attribute_filter production.
	ExitAttribute_filter(c *Attribute_filterContext)

	// ExitBind_rule is called when exiting the bind_rule production.
	ExitBind_rule(c *Bind_ruleContext)

	// ExitParenthetical_bind_rule is called when exiting the parenthetical_bind_rule production.
	ExitParenthetical_bind_rule(c *Parenthetical_bind_ruleContext)

	// ExitParenthetical_bind_rule_req_bool_op is called when exiting the parenthetical_bind_rule_req_bool_op production.
	ExitParenthetical_bind_rule_req_bool_op(c *Parenthetical_bind_rule_req_bool_opContext)

	// ExitNegated_bind_rule_expression is called when exiting the negated_bind_rule_expression production.
	ExitNegated_bind_rule_expression(c *Negated_bind_rule_expressionContext)

	// ExitParenthetical_bind_rule_expression is called when exiting the parenthetical_bind_rule_expression production.
	ExitParenthetical_bind_rule_expression(c *Parenthetical_bind_rule_expressionContext)

	// ExitBind_rule_expression_recursion is called when exiting the bind_rule_expression_recursion production.
	ExitBind_rule_expression_recursion(c *Bind_rule_expression_recursionContext)

	// ExitBind_rule_expression is called when exiting the bind_rule_expression production.
	ExitBind_rule_expression(c *Bind_rule_expressionContext)

	// ExitUserdn_expression is called when exiting the userdn_expression production.
	ExitUserdn_expression(c *Userdn_expressionContext)

	// ExitUserattr_expression is called when exiting the userattr_expression production.
	ExitUserattr_expression(c *Userattr_expressionContext)

	// ExitGroupdn_expression is called when exiting the groupdn_expression production.
	ExitGroupdn_expression(c *Groupdn_expressionContext)

	// ExitGroupattr_expression is called when exiting the groupattr_expression production.
	ExitGroupattr_expression(c *Groupattr_expressionContext)

	// ExitRoledn_expression is called when exiting the roledn_expression production.
	ExitRoledn_expression(c *Roledn_expressionContext)

	// ExitDns_expression is called when exiting the dns_expression production.
	ExitDns_expression(c *Dns_expressionContext)

	// ExitIp_expression is called when exiting the ip_expression production.
	ExitIp_expression(c *Ip_expressionContext)

	// ExitTimeofday_expression is called when exiting the timeofday_expression production.
	ExitTimeofday_expression(c *Timeofday_expressionContext)

	// ExitDayofweek_expression is called when exiting the dayofweek_expression production.
	ExitDayofweek_expression(c *Dayofweek_expressionContext)

	// ExitSsf_expression is called when exiting the ssf_expression production.
	ExitSsf_expression(c *Ssf_expressionContext)

	// ExitAuthmethod_expression is called when exiting the authmethod_expression production.
	ExitAuthmethod_expression(c *Authmethod_expressionContext)

	// ExitParenthetical_dayofweek_bind_rule is called when exiting the parenthetical_dayofweek_bind_rule production.
	ExitParenthetical_dayofweek_bind_rule(c *Parenthetical_dayofweek_bind_ruleContext)

	// ExitDayofweek_bind_rule is called when exiting the dayofweek_bind_rule production.
	ExitDayofweek_bind_rule(c *Dayofweek_bind_ruleContext)

	// ExitSun is called when exiting the Sun production.
	ExitSun(c *SunContext)

	// ExitMon is called when exiting the Mon production.
	ExitMon(c *MonContext)

	// ExitTues is called when exiting the Tues production.
	ExitTues(c *TuesContext)

	// ExitWed is called when exiting the Wed production.
	ExitWed(c *WedContext)

	// ExitThur is called when exiting the Thur production.
	ExitThur(c *ThurContext)

	// ExitFri is called when exiting the Fri production.
	ExitFri(c *FriContext)

	// ExitSat is called when exiting the Sat production.
	ExitSat(c *SatContext)

	// ExitParentheticalAuthenticationMethod is called when exiting the parentheticalAuthenticationMethod production.
	ExitParentheticalAuthenticationMethod(c *ParentheticalAuthenticationMethodContext)

	// ExitAuthentication_method is called when exiting the authentication_method production.
	ExitAuthentication_method(c *Authentication_methodContext)

	// ExitNone is called when exiting the none production.
	ExitNone(c *NoneContext)

	// ExitSimple is called when exiting the simple production.
	ExitSimple(c *SimpleContext)

	// ExitSsl is called when exiting the ssl production.
	ExitSsl(c *SslContext)

	// ExitSasl is called when exiting the sasl production.
	ExitSasl(c *SaslContext)

	// ExitParenthetical_bind_userdn is called when exiting the parenthetical_bind_userdn production.
	ExitParenthetical_bind_userdn(c *Parenthetical_bind_userdnContext)

	// ExitBind_userdn is called when exiting the bind_userdn production.
	ExitBind_userdn(c *Bind_userdnContext)

	// ExitParenthetical_bind_roledn is called when exiting the parenthetical_bind_roledn production.
	ExitParenthetical_bind_roledn(c *Parenthetical_bind_rolednContext)

	// ExitBind_roledn is called when exiting the bind_roledn production.
	ExitBind_roledn(c *Bind_rolednContext)

	// ExitParenthetical_bind_groupdn is called when exiting the parenthetical_bind_groupdn production.
	ExitParenthetical_bind_groupdn(c *Parenthetical_bind_groupdnContext)

	// ExitBind_groupdn is called when exiting the bind_groupdn production.
	ExitBind_groupdn(c *Bind_groupdnContext)

	// ExitParenthetical_bind_userattr is called when exiting the parenthetical_bind_userattr production.
	ExitParenthetical_bind_userattr(c *Parenthetical_bind_userattrContext)

	// ExitBind_userattr is called when exiting the bind_userattr production.
	ExitBind_userattr(c *Bind_userattrContext)

	// ExitParenthetical_bind_groupattr is called when exiting the parenthetical_bind_groupattr production.
	ExitParenthetical_bind_groupattr(c *Parenthetical_bind_groupattrContext)

	// ExitBind_groupattr is called when exiting the bind_groupattr production.
	ExitBind_groupattr(c *Bind_groupattrContext)

	// ExitParenthetical_ssf is called when exiting the parenthetical_ssf production.
	ExitParenthetical_ssf(c *Parenthetical_ssfContext)

	// ExitBind_ssf is called when exiting the bind_ssf production.
	ExitBind_ssf(c *Bind_ssfContext)

	// ExitParenthetical_bind_timeofday is called when exiting the parenthetical_bind_timeofday production.
	ExitParenthetical_bind_timeofday(c *Parenthetical_bind_timeofdayContext)

	// ExitBind_timeofday is called when exiting the bind_timeofday production.
	ExitBind_timeofday(c *Bind_timeofdayContext)

	// ExitParenthetical_bind_ip is called when exiting the parenthetical_bind_ip production.
	ExitParenthetical_bind_ip(c *Parenthetical_bind_ipContext)

	// ExitBind_ip is called when exiting the bind_ip production.
	ExitBind_ip(c *Bind_ipContext)

	// ExitParenthetical_bind_dns is called when exiting the parenthetical_bind_dns production.
	ExitParenthetical_bind_dns(c *Parenthetical_bind_dnsContext)

	// ExitDns_bind_rule is called when exiting the dns_bind_rule production.
	ExitDns_bind_rule(c *Dns_bind_ruleContext)

	// ExitIps is called when exiting the ips production.
	ExitIps(c *IpsContext)

	// ExitIpv4_address is called when exiting the ipv4_address production.
	ExitIpv4_address(c *Ipv4_addressContext)

	// ExitIpv6_address is called when exiting the ipv6_address production.
	ExitIpv6_address(c *Ipv6_addressContext)

	// ExitIpv4 is called when exiting the ipv4 production.
	ExitIpv4(c *Ipv4Context)

	// ExitIpv6 is called when exiting the ipv6 production.
	ExitIpv6(c *Ipv6Context)

	// ExitFqdn is called when exiting the fqdn production.
	ExitFqdn(c *FqdnContext)

	// ExitFullyQualifiedLDAPURI is called when exiting the fullyQualifiedLDAPURI production.
	ExitFullyQualifiedLDAPURI(c *FullyQualifiedLDAPURIContext)

	// ExitUriSearchFilter is called when exiting the uriSearchFilter production.
	ExitUriSearchFilter(c *UriSearchFilterContext)

	// ExitUriSearchScopes is called when exiting the uriSearchScopes production.
	ExitUriSearchScopes(c *UriSearchScopesContext)

	// ExitUriAttributeList is called when exiting the uriAttributeList production.
	ExitUriAttributeList(c *UriAttributeListContext)

	// ExitQuoted_distinguished_name_list is called when exiting the quoted_distinguished_name_list production.
	ExitQuoted_distinguished_name_list(c *Quoted_distinguished_name_listContext)

	// ExitList_of_quoted_distinguished_names is called when exiting the list_of_quoted_distinguished_names production.
	ExitList_of_quoted_distinguished_names(c *List_of_quoted_distinguished_namesContext)

	// ExitDn is called when exiting the dn production.
	ExitDn(c *DnContext)

	// ExitAnonymous_dn_alias is called when exiting the anonymous_dn_alias production.
	ExitAnonymous_dn_alias(c *Anonymous_dn_aliasContext)

	// ExitAny_user_dn_alias is called when exiting the any_user_dn_alias production.
	ExitAny_user_dn_alias(c *Any_user_dn_aliasContext)

	// ExitSelf_dn_alias is called when exiting the self_dn_alias production.
	ExitSelf_dn_alias(c *Self_dn_aliasContext)

	// ExitParent_dn_alias is called when exiting the parent_dn_alias production.
	ExitParent_dn_alias(c *Parent_dn_aliasContext)

	// ExitRdn_macro is called when exiting the rdn_macro production.
	ExitRdn_macro(c *Rdn_macroContext)

	// ExitParenthetical_filter_expression is called when exiting the parenthetical_filter_expression production.
	ExitParenthetical_filter_expression(c *Parenthetical_filter_expressionContext)

	// ExitFilter_expressions is called when exiting the filter_expressions production.
	ExitFilter_expressions(c *Filter_expressionsContext)

	// ExitParenthetical_filter_expression_opt_bool is called when exiting the parenthetical_filter_expression_opt_bool production.
	ExitParenthetical_filter_expression_opt_bool(c *Parenthetical_filter_expression_opt_boolContext)

	// ExitNot_filter_expression is called when exiting the not_filter_expression production.
	ExitNot_filter_expression(c *Not_filter_expressionContext)

	// ExitAva_expression is called when exiting the ava_expression production.
	ExitAva_expression(c *Ava_expressionContext)

	// ExitAva_or_rdn is called when exiting the ava_or_rdn production.
	ExitAva_or_rdn(c *Ava_or_rdnContext)

	// ExitInheritance_expression is called when exiting the inheritance_expression production.
	ExitInheritance_expression(c *Inheritance_expressionContext)

	// ExitInheritance_levels is called when exiting the inheritance_levels production.
	ExitInheritance_levels(c *Inheritance_levelsContext)

	// ExitAttr_bind_type_or_value is called when exiting the attr_bind_type_or_value production.
	ExitAttr_bind_type_or_value(c *Attr_bind_type_or_valueContext)

	// ExitUSERDN is called when exiting the USERDN production.
	ExitUSERDN(c *USERDNContext)

	// ExitGROUPDN is called when exiting the GROUPDN production.
	ExitGROUPDN(c *GROUPDNContext)

	// ExitROLEDN is called when exiting the ROLEDN production.
	ExitROLEDN(c *ROLEDNContext)

	// ExitSELFDN is called when exiting the SELFDN production.
	ExitSELFDN(c *SELFDNContext)

	// ExitLDAPURL is called when exiting the LDAPURL production.
	ExitLDAPURL(c *LDAPURLContext)

	// ExitKey_or_value is called when exiting the key_or_value production.
	ExitKey_or_value(c *Key_or_valueContext)

	// ExitEqual_to is called when exiting the equal_to production.
	ExitEqual_to(c *Equal_toContext)

	// ExitNot_equal_to is called when exiting the not_equal_to production.
	ExitNot_equal_to(c *Not_equal_toContext)

	// ExitGreater_than_or_equal is called when exiting the greater_than_or_equal production.
	ExitGreater_than_or_equal(c *Greater_than_or_equalContext)

	// ExitLess_than_or_equal is called when exiting the less_than_or_equal production.
	ExitLess_than_or_equal(c *Less_than_or_equalContext)

	// ExitApprox is called when exiting the approx production.
	ExitApprox(c *ApproxContext)

	// ExitExtensible_rule is called when exiting the extensible_rule production.
	ExitExtensible_rule(c *Extensible_ruleContext)

	// ExitExtensible_rule_with_dn is called when exiting the extensible_rule_with_dn production.
	ExitExtensible_rule_with_dn(c *Extensible_rule_with_dnContext)

	// ExitExtensible_rule_with_attrs is called when exiting the extensible_rule_with_attrs production.
	ExitExtensible_rule_with_attrs(c *Extensible_rule_with_attrsContext)

	// ExitExtensible_rule_with_dn_oid is called when exiting the extensible_rule_with_dn_oid production.
	ExitExtensible_rule_with_dn_oid(c *Extensible_rule_with_dn_oidContext)

	// ExitEqualTo is called when exiting the equalTo production.
	ExitEqualTo(c *EqualToContext)

	// ExitNotEqualTo is called when exiting the notEqualTo production.
	ExitNotEqualTo(c *NotEqualToContext)

	// ExitGreaterThan is called when exiting the greaterThan production.
	ExitGreaterThan(c *GreaterThanContext)

	// ExitLessThan is called when exiting the lessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitGreaterThanOrEqual is called when exiting the greaterThanOrEqual production.
	ExitGreaterThanOrEqual(c *GreaterThanOrEqualContext)

	// ExitLessThanOrEqual is called when exiting the lessThanOrEqual production.
	ExitLessThanOrEqual(c *LessThanOrEqualContext)

	// ExitApproximate is called when exiting the approximate production.
	ExitApproximate(c *ApproximateContext)

	// ExitExtensibleRule is called when exiting the extensibleRule production.
	ExitExtensibleRule(c *ExtensibleRuleContext)

	// ExitExtensibleRuleDNOID is called when exiting the extensibleRuleDNOID production.
	ExitExtensibleRuleDNOID(c *ExtensibleRuleDNOIDContext)

	// ExitExtensibleRuleDN is called when exiting the extensibleRuleDN production.
	ExitExtensibleRuleDN(c *ExtensibleRuleDNContext)

	// ExitExtensibleRuleAttr is called when exiting the extensibleRuleAttr production.
	ExitExtensibleRuleAttr(c *ExtensibleRuleAttrContext)

	// ExitORDelimiter is called when exiting the oRDelimiter production.
	ExitORDelimiter(c *ORDelimiterContext)

	// ExitANDDelimiter is called when exiting the aNDDelimiter production.
	ExitANDDelimiter(c *ANDDelimiterContext)
}
