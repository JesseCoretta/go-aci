// Code generated from ACIParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ACIParser
import "github.com/antlr4-go/antlr/v4"

// BaseACIParserListener is a complete listener for a parse tree produced by ACIParser.
type BaseACIParserListener struct{}

var _ ACIParserListener = &BaseACIParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseACIParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseACIParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseACIParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseACIParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseACIParserListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseACIParserListener) ExitParse(ctx *ParseContext) {}

// EnterAci is called when production aci is entered.
func (s *BaseACIParserListener) EnterAci(ctx *AciContext) {}

// ExitAci is called when production aci is exited.
func (s *BaseACIParserListener) ExitAci(ctx *AciContext) {}

// EnterPermission_bind_rules is called when production permission_bind_rules is entered.
func (s *BaseACIParserListener) EnterPermission_bind_rules(ctx *Permission_bind_rulesContext) {}

// ExitPermission_bind_rules is called when production permission_bind_rules is exited.
func (s *BaseACIParserListener) ExitPermission_bind_rules(ctx *Permission_bind_rulesContext) {}

// EnterPermission_and_bind_rule_pair is called when production permission_and_bind_rule_pair is entered.
func (s *BaseACIParserListener) EnterPermission_and_bind_rule_pair(ctx *Permission_and_bind_rule_pairContext) {
}

// ExitPermission_and_bind_rule_pair is called when production permission_and_bind_rule_pair is exited.
func (s *BaseACIParserListener) ExitPermission_and_bind_rule_pair(ctx *Permission_and_bind_rule_pairContext) {
}

// EnterPermission_expression is called when production permission_expression is entered.
func (s *BaseACIParserListener) EnterPermission_expression(ctx *Permission_expressionContext) {}

// ExitPermission_expression is called when production permission_expression is exited.
func (s *BaseACIParserListener) ExitPermission_expression(ctx *Permission_expressionContext) {}

// EnterAllow_access is called when production allow_access is entered.
func (s *BaseACIParserListener) EnterAllow_access(ctx *Allow_accessContext) {}

// ExitAllow_access is called when production allow_access is exited.
func (s *BaseACIParserListener) ExitAllow_access(ctx *Allow_accessContext) {}

// EnterDeny_access is called when production deny_access is entered.
func (s *BaseACIParserListener) EnterDeny_access(ctx *Deny_accessContext) {}

// ExitDeny_access is called when production deny_access is exited.
func (s *BaseACIParserListener) ExitDeny_access(ctx *Deny_accessContext) {}

// EnterSearch_privilege is called when production search_privilege is entered.
func (s *BaseACIParserListener) EnterSearch_privilege(ctx *Search_privilegeContext) {}

// ExitSearch_privilege is called when production search_privilege is exited.
func (s *BaseACIParserListener) ExitSearch_privilege(ctx *Search_privilegeContext) {}

// EnterRead_privilege is called when production read_privilege is entered.
func (s *BaseACIParserListener) EnterRead_privilege(ctx *Read_privilegeContext) {}

// ExitRead_privilege is called when production read_privilege is exited.
func (s *BaseACIParserListener) ExitRead_privilege(ctx *Read_privilegeContext) {}

// EnterCompare_privilege is called when production compare_privilege is entered.
func (s *BaseACIParserListener) EnterCompare_privilege(ctx *Compare_privilegeContext) {}

// ExitCompare_privilege is called when production compare_privilege is exited.
func (s *BaseACIParserListener) ExitCompare_privilege(ctx *Compare_privilegeContext) {}

// EnterAdd_privilege is called when production add_privilege is entered.
func (s *BaseACIParserListener) EnterAdd_privilege(ctx *Add_privilegeContext) {}

// ExitAdd_privilege is called when production add_privilege is exited.
func (s *BaseACIParserListener) ExitAdd_privilege(ctx *Add_privilegeContext) {}

// EnterDelete_privilege is called when production delete_privilege is entered.
func (s *BaseACIParserListener) EnterDelete_privilege(ctx *Delete_privilegeContext) {}

// ExitDelete_privilege is called when production delete_privilege is exited.
func (s *BaseACIParserListener) ExitDelete_privilege(ctx *Delete_privilegeContext) {}

// EnterSelfwrite_privilege is called when production selfwrite_privilege is entered.
func (s *BaseACIParserListener) EnterSelfwrite_privilege(ctx *Selfwrite_privilegeContext) {}

// ExitSelfwrite_privilege is called when production selfwrite_privilege is exited.
func (s *BaseACIParserListener) ExitSelfwrite_privilege(ctx *Selfwrite_privilegeContext) {}

// EnterProxy_privilege is called when production proxy_privilege is entered.
func (s *BaseACIParserListener) EnterProxy_privilege(ctx *Proxy_privilegeContext) {}

// ExitProxy_privilege is called when production proxy_privilege is exited.
func (s *BaseACIParserListener) ExitProxy_privilege(ctx *Proxy_privilegeContext) {}

// EnterImport_privilege is called when production import_privilege is entered.
func (s *BaseACIParserListener) EnterImport_privilege(ctx *Import_privilegeContext) {}

// ExitImport_privilege is called when production import_privilege is exited.
func (s *BaseACIParserListener) ExitImport_privilege(ctx *Import_privilegeContext) {}

// EnterExport_privilege is called when production export_privilege is entered.
func (s *BaseACIParserListener) EnterExport_privilege(ctx *Export_privilegeContext) {}

// ExitExport_privilege is called when production export_privilege is exited.
func (s *BaseACIParserListener) ExitExport_privilege(ctx *Export_privilegeContext) {}

// EnterAll_privileges is called when production all_privileges is entered.
func (s *BaseACIParserListener) EnterAll_privileges(ctx *All_privilegesContext) {}

// ExitAll_privileges is called when production all_privileges is exited.
func (s *BaseACIParserListener) ExitAll_privileges(ctx *All_privilegesContext) {}

// EnterTarget_rules is called when production target_rules is entered.
func (s *BaseACIParserListener) EnterTarget_rules(ctx *Target_rulesContext) {}

// ExitTarget_rules is called when production target_rules is exited.
func (s *BaseACIParserListener) ExitTarget_rules(ctx *Target_rulesContext) {}

// EnterTargetRule is called when production targetRule is entered.
func (s *BaseACIParserListener) EnterTargetRule(ctx *TargetRuleContext) {}

// ExitTargetRule is called when production targetRule is exited.
func (s *BaseACIParserListener) ExitTargetRule(ctx *TargetRuleContext) {}

// EnterTarget_dn_rule is called when production target_dn_rule is entered.
func (s *BaseACIParserListener) EnterTarget_dn_rule(ctx *Target_dn_ruleContext) {}

// ExitTarget_dn_rule is called when production target_dn_rule is exited.
func (s *BaseACIParserListener) ExitTarget_dn_rule(ctx *Target_dn_ruleContext) {}

// EnterTarget_to_rule is called when production target_to_rule is entered.
func (s *BaseACIParserListener) EnterTarget_to_rule(ctx *Target_to_ruleContext) {}

// ExitTarget_to_rule is called when production target_to_rule is exited.
func (s *BaseACIParserListener) ExitTarget_to_rule(ctx *Target_to_ruleContext) {}

// EnterTarget_from_rule is called when production target_from_rule is entered.
func (s *BaseACIParserListener) EnterTarget_from_rule(ctx *Target_from_ruleContext) {}

// ExitTarget_from_rule is called when production target_from_rule is exited.
func (s *BaseACIParserListener) ExitTarget_from_rule(ctx *Target_from_ruleContext) {}

// EnterTargetfilter_rule is called when production targetfilter_rule is entered.
func (s *BaseACIParserListener) EnterTargetfilter_rule(ctx *Targetfilter_ruleContext) {}

// ExitTargetfilter_rule is called when production targetfilter_rule is exited.
func (s *BaseACIParserListener) ExitTargetfilter_rule(ctx *Targetfilter_ruleContext) {}

// EnterTargattrfilters_rule is called when production targattrfilters_rule is entered.
func (s *BaseACIParserListener) EnterTargattrfilters_rule(ctx *Targattrfilters_ruleContext) {}

// ExitTargattrfilters_rule is called when production targattrfilters_rule is exited.
func (s *BaseACIParserListener) ExitTargattrfilters_rule(ctx *Targattrfilters_ruleContext) {}

// EnterTargetscope_rule is called when production targetscope_rule is entered.
func (s *BaseACIParserListener) EnterTargetscope_rule(ctx *Targetscope_ruleContext) {}

// ExitTargetscope_rule is called when production targetscope_rule is exited.
func (s *BaseACIParserListener) ExitTargetscope_rule(ctx *Targetscope_ruleContext) {}

// EnterTargetattr_rule is called when production targetattr_rule is entered.
func (s *BaseACIParserListener) EnterTargetattr_rule(ctx *Targetattr_ruleContext) {}

// ExitTargetattr_rule is called when production targetattr_rule is exited.
func (s *BaseACIParserListener) ExitTargetattr_rule(ctx *Targetattr_ruleContext) {}

// EnterTargetcontrol_rule is called when production targetcontrol_rule is entered.
func (s *BaseACIParserListener) EnterTargetcontrol_rule(ctx *Targetcontrol_ruleContext) {}

// ExitTargetcontrol_rule is called when production targetcontrol_rule is exited.
func (s *BaseACIParserListener) ExitTargetcontrol_rule(ctx *Targetcontrol_ruleContext) {}

// EnterTargetextop_rule is called when production targetextop_rule is entered.
func (s *BaseACIParserListener) EnterTargetextop_rule(ctx *Targetextop_ruleContext) {}

// ExitTargetextop_rule is called when production targetextop_rule is exited.
func (s *BaseACIParserListener) ExitTargetextop_rule(ctx *Targetextop_ruleContext) {}

// EnterBase_object_targetscope is called when production base_object_targetscope is entered.
func (s *BaseACIParserListener) EnterBase_object_targetscope(ctx *Base_object_targetscopeContext) {}

// ExitBase_object_targetscope is called when production base_object_targetscope is exited.
func (s *BaseACIParserListener) ExitBase_object_targetscope(ctx *Base_object_targetscopeContext) {}

// EnterOne_level_targetscope is called when production one_level_targetscope is entered.
func (s *BaseACIParserListener) EnterOne_level_targetscope(ctx *One_level_targetscopeContext) {}

// ExitOne_level_targetscope is called when production one_level_targetscope is exited.
func (s *BaseACIParserListener) ExitOne_level_targetscope(ctx *One_level_targetscopeContext) {}

// EnterSub_tree_targetscope is called when production sub_tree_targetscope is entered.
func (s *BaseACIParserListener) EnterSub_tree_targetscope(ctx *Sub_tree_targetscopeContext) {}

// ExitSub_tree_targetscope is called when production sub_tree_targetscope is exited.
func (s *BaseACIParserListener) ExitSub_tree_targetscope(ctx *Sub_tree_targetscopeContext) {}

// EnterSubordinate_targetscope is called when production subordinate_targetscope is entered.
func (s *BaseACIParserListener) EnterSubordinate_targetscope(ctx *Subordinate_targetscopeContext) {}

// ExitSubordinate_targetscope is called when production subordinate_targetscope is exited.
func (s *BaseACIParserListener) ExitSubordinate_targetscope(ctx *Subordinate_targetscopeContext) {}

// EnterQuoted_object_identifier_list is called when production quoted_object_identifier_list is entered.
func (s *BaseACIParserListener) EnterQuoted_object_identifier_list(ctx *Quoted_object_identifier_listContext) {
}

// ExitQuoted_object_identifier_list is called when production quoted_object_identifier_list is exited.
func (s *BaseACIParserListener) ExitQuoted_object_identifier_list(ctx *Quoted_object_identifier_listContext) {
}

// EnterList_of_quoted_object_identifiers is called when production list_of_quoted_object_identifiers is entered.
func (s *BaseACIParserListener) EnterList_of_quoted_object_identifiers(ctx *List_of_quoted_object_identifiersContext) {
}

// ExitList_of_quoted_object_identifiers is called when production list_of_quoted_object_identifiers is exited.
func (s *BaseACIParserListener) ExitList_of_quoted_object_identifiers(ctx *List_of_quoted_object_identifiersContext) {
}

// EnterQuoted_targeted_attributes_list is called when production quoted_targeted_attributes_list is entered.
func (s *BaseACIParserListener) EnterQuoted_targeted_attributes_list(ctx *Quoted_targeted_attributes_listContext) {
}

// ExitQuoted_targeted_attributes_list is called when production quoted_targeted_attributes_list is exited.
func (s *BaseACIParserListener) ExitQuoted_targeted_attributes_list(ctx *Quoted_targeted_attributes_listContext) {
}

// EnterList_of_quoted_attributes is called when production list_of_quoted_attributes is entered.
func (s *BaseACIParserListener) EnterList_of_quoted_attributes(ctx *List_of_quoted_attributesContext) {
}

// ExitList_of_quoted_attributes is called when production list_of_quoted_attributes is exited.
func (s *BaseACIParserListener) ExitList_of_quoted_attributes(ctx *List_of_quoted_attributesContext) {
}

// EnterObject_identifier is called when production object_identifier is entered.
func (s *BaseACIParserListener) EnterObject_identifier(ctx *Object_identifierContext) {}

// ExitObject_identifier is called when production object_identifier is exited.
func (s *BaseACIParserListener) ExitObject_identifier(ctx *Object_identifierContext) {}

// EnterNumber_form is called when production number_form is entered.
func (s *BaseACIParserListener) EnterNumber_form(ctx *Number_formContext) {}

// ExitNumber_form is called when production number_form is exited.
func (s *BaseACIParserListener) ExitNumber_form(ctx *Number_formContext) {}

// EnterAttribute_filters_sets is called when production attribute_filters_sets is entered.
func (s *BaseACIParserListener) EnterAttribute_filters_sets(ctx *Attribute_filters_setsContext) {}

// ExitAttribute_filters_sets is called when production attribute_filters_sets is exited.
func (s *BaseACIParserListener) ExitAttribute_filters_sets(ctx *Attribute_filters_setsContext) {}

// EnterAttribute_filters_set is called when production attribute_filters_set is entered.
func (s *BaseACIParserListener) EnterAttribute_filters_set(ctx *Attribute_filters_setContext) {}

// ExitAttribute_filters_set is called when production attribute_filters_set is exited.
func (s *BaseACIParserListener) ExitAttribute_filters_set(ctx *Attribute_filters_setContext) {}

// EnterAttribute_filter_single is called when production attribute_filter_single is entered.
func (s *BaseACIParserListener) EnterAttribute_filter_single(ctx *Attribute_filter_singleContext) {}

// ExitAttribute_filter_single is called when production attribute_filter_single is exited.
func (s *BaseACIParserListener) ExitAttribute_filter_single(ctx *Attribute_filter_singleContext) {}

// EnterAttribute_filters is called when production attribute_filters is entered.
func (s *BaseACIParserListener) EnterAttribute_filters(ctx *Attribute_filtersContext) {}

// ExitAttribute_filters is called when production attribute_filters is exited.
func (s *BaseACIParserListener) ExitAttribute_filters(ctx *Attribute_filtersContext) {}

// EnterAttribute_filter_set is called when production attribute_filter_set is entered.
func (s *BaseACIParserListener) EnterAttribute_filter_set(ctx *Attribute_filter_setContext) {}

// ExitAttribute_filter_set is called when production attribute_filter_set is exited.
func (s *BaseACIParserListener) ExitAttribute_filter_set(ctx *Attribute_filter_setContext) {}

// EnterAdd_filter_operation is called when production add_filter_operation is entered.
func (s *BaseACIParserListener) EnterAdd_filter_operation(ctx *Add_filter_operationContext) {}

// ExitAdd_filter_operation is called when production add_filter_operation is exited.
func (s *BaseACIParserListener) ExitAdd_filter_operation(ctx *Add_filter_operationContext) {}

// EnterDelete_filter_operation is called when production delete_filter_operation is entered.
func (s *BaseACIParserListener) EnterDelete_filter_operation(ctx *Delete_filter_operationContext) {}

// ExitDelete_filter_operation is called when production delete_filter_operation is exited.
func (s *BaseACIParserListener) ExitDelete_filter_operation(ctx *Delete_filter_operationContext) {}

// EnterAttribute_filter is called when production attribute_filter is entered.
func (s *BaseACIParserListener) EnterAttribute_filter(ctx *Attribute_filterContext) {}

// ExitAttribute_filter is called when production attribute_filter is exited.
func (s *BaseACIParserListener) ExitAttribute_filter(ctx *Attribute_filterContext) {}

// EnterBind_rule is called when production bind_rule is entered.
func (s *BaseACIParserListener) EnterBind_rule(ctx *Bind_ruleContext) {}

// ExitBind_rule is called when production bind_rule is exited.
func (s *BaseACIParserListener) ExitBind_rule(ctx *Bind_ruleContext) {}

// EnterParenthetical_bind_rule is called when production parenthetical_bind_rule is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_rule(ctx *Parenthetical_bind_ruleContext) {}

// ExitParenthetical_bind_rule is called when production parenthetical_bind_rule is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_rule(ctx *Parenthetical_bind_ruleContext) {}

// EnterParenthetical_bind_rule_req_bool_op is called when production parenthetical_bind_rule_req_bool_op is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_rule_req_bool_op(ctx *Parenthetical_bind_rule_req_bool_opContext) {
}

// ExitParenthetical_bind_rule_req_bool_op is called when production parenthetical_bind_rule_req_bool_op is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_rule_req_bool_op(ctx *Parenthetical_bind_rule_req_bool_opContext) {
}

// EnterNegated_bind_rule_expression is called when production negated_bind_rule_expression is entered.
func (s *BaseACIParserListener) EnterNegated_bind_rule_expression(ctx *Negated_bind_rule_expressionContext) {
}

// ExitNegated_bind_rule_expression is called when production negated_bind_rule_expression is exited.
func (s *BaseACIParserListener) ExitNegated_bind_rule_expression(ctx *Negated_bind_rule_expressionContext) {
}

// EnterParenthetical_bind_rule_expression is called when production parenthetical_bind_rule_expression is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_rule_expression(ctx *Parenthetical_bind_rule_expressionContext) {
}

// ExitParenthetical_bind_rule_expression is called when production parenthetical_bind_rule_expression is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_rule_expression(ctx *Parenthetical_bind_rule_expressionContext) {
}

// EnterBind_rule_expression_recursion is called when production bind_rule_expression_recursion is entered.
func (s *BaseACIParserListener) EnterBind_rule_expression_recursion(ctx *Bind_rule_expression_recursionContext) {
}

// ExitBind_rule_expression_recursion is called when production bind_rule_expression_recursion is exited.
func (s *BaseACIParserListener) ExitBind_rule_expression_recursion(ctx *Bind_rule_expression_recursionContext) {
}

// EnterBind_rule_expression is called when production bind_rule_expression is entered.
func (s *BaseACIParserListener) EnterBind_rule_expression(ctx *Bind_rule_expressionContext) {}

// ExitBind_rule_expression is called when production bind_rule_expression is exited.
func (s *BaseACIParserListener) ExitBind_rule_expression(ctx *Bind_rule_expressionContext) {}

// EnterUserdn_expression is called when production userdn_expression is entered.
func (s *BaseACIParserListener) EnterUserdn_expression(ctx *Userdn_expressionContext) {}

// ExitUserdn_expression is called when production userdn_expression is exited.
func (s *BaseACIParserListener) ExitUserdn_expression(ctx *Userdn_expressionContext) {}

// EnterUserattr_expression is called when production userattr_expression is entered.
func (s *BaseACIParserListener) EnterUserattr_expression(ctx *Userattr_expressionContext) {}

// ExitUserattr_expression is called when production userattr_expression is exited.
func (s *BaseACIParserListener) ExitUserattr_expression(ctx *Userattr_expressionContext) {}

// EnterGroupdn_expression is called when production groupdn_expression is entered.
func (s *BaseACIParserListener) EnterGroupdn_expression(ctx *Groupdn_expressionContext) {}

// ExitGroupdn_expression is called when production groupdn_expression is exited.
func (s *BaseACIParserListener) ExitGroupdn_expression(ctx *Groupdn_expressionContext) {}

// EnterGroupattr_expression is called when production groupattr_expression is entered.
func (s *BaseACIParserListener) EnterGroupattr_expression(ctx *Groupattr_expressionContext) {}

// ExitGroupattr_expression is called when production groupattr_expression is exited.
func (s *BaseACIParserListener) ExitGroupattr_expression(ctx *Groupattr_expressionContext) {}

// EnterRoledn_expression is called when production roledn_expression is entered.
func (s *BaseACIParserListener) EnterRoledn_expression(ctx *Roledn_expressionContext) {}

// ExitRoledn_expression is called when production roledn_expression is exited.
func (s *BaseACIParserListener) ExitRoledn_expression(ctx *Roledn_expressionContext) {}

// EnterDns_expression is called when production dns_expression is entered.
func (s *BaseACIParserListener) EnterDns_expression(ctx *Dns_expressionContext) {}

// ExitDns_expression is called when production dns_expression is exited.
func (s *BaseACIParserListener) ExitDns_expression(ctx *Dns_expressionContext) {}

// EnterIp_expression is called when production ip_expression is entered.
func (s *BaseACIParserListener) EnterIp_expression(ctx *Ip_expressionContext) {}

// ExitIp_expression is called when production ip_expression is exited.
func (s *BaseACIParserListener) ExitIp_expression(ctx *Ip_expressionContext) {}

// EnterTimeofday_expression is called when production timeofday_expression is entered.
func (s *BaseACIParserListener) EnterTimeofday_expression(ctx *Timeofday_expressionContext) {}

// ExitTimeofday_expression is called when production timeofday_expression is exited.
func (s *BaseACIParserListener) ExitTimeofday_expression(ctx *Timeofday_expressionContext) {}

// EnterDayofweek_expression is called when production dayofweek_expression is entered.
func (s *BaseACIParserListener) EnterDayofweek_expression(ctx *Dayofweek_expressionContext) {}

// ExitDayofweek_expression is called when production dayofweek_expression is exited.
func (s *BaseACIParserListener) ExitDayofweek_expression(ctx *Dayofweek_expressionContext) {}

// EnterSsf_expression is called when production ssf_expression is entered.
func (s *BaseACIParserListener) EnterSsf_expression(ctx *Ssf_expressionContext) {}

// ExitSsf_expression is called when production ssf_expression is exited.
func (s *BaseACIParserListener) ExitSsf_expression(ctx *Ssf_expressionContext) {}

// EnterAuthmethod_expression is called when production authmethod_expression is entered.
func (s *BaseACIParserListener) EnterAuthmethod_expression(ctx *Authmethod_expressionContext) {}

// ExitAuthmethod_expression is called when production authmethod_expression is exited.
func (s *BaseACIParserListener) ExitAuthmethod_expression(ctx *Authmethod_expressionContext) {}

// EnterParenthetical_dayofweek_bind_rule is called when production parenthetical_dayofweek_bind_rule is entered.
func (s *BaseACIParserListener) EnterParenthetical_dayofweek_bind_rule(ctx *Parenthetical_dayofweek_bind_ruleContext) {
}

// ExitParenthetical_dayofweek_bind_rule is called when production parenthetical_dayofweek_bind_rule is exited.
func (s *BaseACIParserListener) ExitParenthetical_dayofweek_bind_rule(ctx *Parenthetical_dayofweek_bind_ruleContext) {
}

// EnterDayofweek_bind_rule is called when production dayofweek_bind_rule is entered.
func (s *BaseACIParserListener) EnterDayofweek_bind_rule(ctx *Dayofweek_bind_ruleContext) {}

// ExitDayofweek_bind_rule is called when production dayofweek_bind_rule is exited.
func (s *BaseACIParserListener) ExitDayofweek_bind_rule(ctx *Dayofweek_bind_ruleContext) {}

// EnterSun is called when production Sun is entered.
func (s *BaseACIParserListener) EnterSun(ctx *SunContext) {}

// ExitSun is called when production Sun is exited.
func (s *BaseACIParserListener) ExitSun(ctx *SunContext) {}

// EnterMon is called when production Mon is entered.
func (s *BaseACIParserListener) EnterMon(ctx *MonContext) {}

// ExitMon is called when production Mon is exited.
func (s *BaseACIParserListener) ExitMon(ctx *MonContext) {}

// EnterTues is called when production Tues is entered.
func (s *BaseACIParserListener) EnterTues(ctx *TuesContext) {}

// ExitTues is called when production Tues is exited.
func (s *BaseACIParserListener) ExitTues(ctx *TuesContext) {}

// EnterWed is called when production Wed is entered.
func (s *BaseACIParserListener) EnterWed(ctx *WedContext) {}

// ExitWed is called when production Wed is exited.
func (s *BaseACIParserListener) ExitWed(ctx *WedContext) {}

// EnterThur is called when production Thur is entered.
func (s *BaseACIParserListener) EnterThur(ctx *ThurContext) {}

// ExitThur is called when production Thur is exited.
func (s *BaseACIParserListener) ExitThur(ctx *ThurContext) {}

// EnterFri is called when production Fri is entered.
func (s *BaseACIParserListener) EnterFri(ctx *FriContext) {}

// ExitFri is called when production Fri is exited.
func (s *BaseACIParserListener) ExitFri(ctx *FriContext) {}

// EnterSat is called when production Sat is entered.
func (s *BaseACIParserListener) EnterSat(ctx *SatContext) {}

// ExitSat is called when production Sat is exited.
func (s *BaseACIParserListener) ExitSat(ctx *SatContext) {}

// EnterParentheticalAuthenticationMethod is called when production parentheticalAuthenticationMethod is entered.
func (s *BaseACIParserListener) EnterParentheticalAuthenticationMethod(ctx *ParentheticalAuthenticationMethodContext) {
}

// ExitParentheticalAuthenticationMethod is called when production parentheticalAuthenticationMethod is exited.
func (s *BaseACIParserListener) ExitParentheticalAuthenticationMethod(ctx *ParentheticalAuthenticationMethodContext) {
}

// EnterAuthentication_method is called when production authentication_method is entered.
func (s *BaseACIParserListener) EnterAuthentication_method(ctx *Authentication_methodContext) {}

// ExitAuthentication_method is called when production authentication_method is exited.
func (s *BaseACIParserListener) ExitAuthentication_method(ctx *Authentication_methodContext) {}

// EnterNone is called when production none is entered.
func (s *BaseACIParserListener) EnterNone(ctx *NoneContext) {}

// ExitNone is called when production none is exited.
func (s *BaseACIParserListener) ExitNone(ctx *NoneContext) {}

// EnterSimple is called when production simple is entered.
func (s *BaseACIParserListener) EnterSimple(ctx *SimpleContext) {}

// ExitSimple is called when production simple is exited.
func (s *BaseACIParserListener) ExitSimple(ctx *SimpleContext) {}

// EnterSsl is called when production ssl is entered.
func (s *BaseACIParserListener) EnterSsl(ctx *SslContext) {}

// ExitSsl is called when production ssl is exited.
func (s *BaseACIParserListener) ExitSsl(ctx *SslContext) {}

// EnterSasl is called when production sasl is entered.
func (s *BaseACIParserListener) EnterSasl(ctx *SaslContext) {}

// ExitSasl is called when production sasl is exited.
func (s *BaseACIParserListener) ExitSasl(ctx *SaslContext) {}

// EnterParenthetical_bind_userdn is called when production parenthetical_bind_userdn is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_userdn(ctx *Parenthetical_bind_userdnContext) {
}

// ExitParenthetical_bind_userdn is called when production parenthetical_bind_userdn is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_userdn(ctx *Parenthetical_bind_userdnContext) {
}

// EnterBind_userdn is called when production bind_userdn is entered.
func (s *BaseACIParserListener) EnterBind_userdn(ctx *Bind_userdnContext) {}

// ExitBind_userdn is called when production bind_userdn is exited.
func (s *BaseACIParserListener) ExitBind_userdn(ctx *Bind_userdnContext) {}

// EnterParenthetical_bind_roledn is called when production parenthetical_bind_roledn is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_roledn(ctx *Parenthetical_bind_rolednContext) {
}

// ExitParenthetical_bind_roledn is called when production parenthetical_bind_roledn is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_roledn(ctx *Parenthetical_bind_rolednContext) {
}

// EnterBind_roledn is called when production bind_roledn is entered.
func (s *BaseACIParserListener) EnterBind_roledn(ctx *Bind_rolednContext) {}

// ExitBind_roledn is called when production bind_roledn is exited.
func (s *BaseACIParserListener) ExitBind_roledn(ctx *Bind_rolednContext) {}

// EnterParenthetical_bind_groupdn is called when production parenthetical_bind_groupdn is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_groupdn(ctx *Parenthetical_bind_groupdnContext) {
}

// ExitParenthetical_bind_groupdn is called when production parenthetical_bind_groupdn is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_groupdn(ctx *Parenthetical_bind_groupdnContext) {
}

// EnterBind_groupdn is called when production bind_groupdn is entered.
func (s *BaseACIParserListener) EnterBind_groupdn(ctx *Bind_groupdnContext) {}

// ExitBind_groupdn is called when production bind_groupdn is exited.
func (s *BaseACIParserListener) ExitBind_groupdn(ctx *Bind_groupdnContext) {}

// EnterParenthetical_bind_userattr is called when production parenthetical_bind_userattr is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_userattr(ctx *Parenthetical_bind_userattrContext) {
}

// ExitParenthetical_bind_userattr is called when production parenthetical_bind_userattr is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_userattr(ctx *Parenthetical_bind_userattrContext) {
}

// EnterBind_userattr is called when production bind_userattr is entered.
func (s *BaseACIParserListener) EnterBind_userattr(ctx *Bind_userattrContext) {}

// ExitBind_userattr is called when production bind_userattr is exited.
func (s *BaseACIParserListener) ExitBind_userattr(ctx *Bind_userattrContext) {}

// EnterParenthetical_bind_groupattr is called when production parenthetical_bind_groupattr is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_groupattr(ctx *Parenthetical_bind_groupattrContext) {
}

// ExitParenthetical_bind_groupattr is called when production parenthetical_bind_groupattr is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_groupattr(ctx *Parenthetical_bind_groupattrContext) {
}

// EnterBind_groupattr is called when production bind_groupattr is entered.
func (s *BaseACIParserListener) EnterBind_groupattr(ctx *Bind_groupattrContext) {}

// ExitBind_groupattr is called when production bind_groupattr is exited.
func (s *BaseACIParserListener) ExitBind_groupattr(ctx *Bind_groupattrContext) {}

// EnterParenthetical_ssf is called when production parenthetical_ssf is entered.
func (s *BaseACIParserListener) EnterParenthetical_ssf(ctx *Parenthetical_ssfContext) {}

// ExitParenthetical_ssf is called when production parenthetical_ssf is exited.
func (s *BaseACIParserListener) ExitParenthetical_ssf(ctx *Parenthetical_ssfContext) {}

// EnterBind_ssf is called when production bind_ssf is entered.
func (s *BaseACIParserListener) EnterBind_ssf(ctx *Bind_ssfContext) {}

// ExitBind_ssf is called when production bind_ssf is exited.
func (s *BaseACIParserListener) ExitBind_ssf(ctx *Bind_ssfContext) {}

// EnterParenthetical_bind_timeofday is called when production parenthetical_bind_timeofday is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_timeofday(ctx *Parenthetical_bind_timeofdayContext) {
}

// ExitParenthetical_bind_timeofday is called when production parenthetical_bind_timeofday is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_timeofday(ctx *Parenthetical_bind_timeofdayContext) {
}

// EnterBind_timeofday is called when production bind_timeofday is entered.
func (s *BaseACIParserListener) EnterBind_timeofday(ctx *Bind_timeofdayContext) {}

// ExitBind_timeofday is called when production bind_timeofday is exited.
func (s *BaseACIParserListener) ExitBind_timeofday(ctx *Bind_timeofdayContext) {}

// EnterParenthetical_bind_ip is called when production parenthetical_bind_ip is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_ip(ctx *Parenthetical_bind_ipContext) {}

// ExitParenthetical_bind_ip is called when production parenthetical_bind_ip is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_ip(ctx *Parenthetical_bind_ipContext) {}

// EnterBind_ip is called when production bind_ip is entered.
func (s *BaseACIParserListener) EnterBind_ip(ctx *Bind_ipContext) {}

// ExitBind_ip is called when production bind_ip is exited.
func (s *BaseACIParserListener) ExitBind_ip(ctx *Bind_ipContext) {}

// EnterParenthetical_bind_dns is called when production parenthetical_bind_dns is entered.
func (s *BaseACIParserListener) EnterParenthetical_bind_dns(ctx *Parenthetical_bind_dnsContext) {}

// ExitParenthetical_bind_dns is called when production parenthetical_bind_dns is exited.
func (s *BaseACIParserListener) ExitParenthetical_bind_dns(ctx *Parenthetical_bind_dnsContext) {}

// EnterDns_bind_rule is called when production dns_bind_rule is entered.
func (s *BaseACIParserListener) EnterDns_bind_rule(ctx *Dns_bind_ruleContext) {}

// ExitDns_bind_rule is called when production dns_bind_rule is exited.
func (s *BaseACIParserListener) ExitDns_bind_rule(ctx *Dns_bind_ruleContext) {}

// EnterIps is called when production ips is entered.
func (s *BaseACIParserListener) EnterIps(ctx *IpsContext) {}

// ExitIps is called when production ips is exited.
func (s *BaseACIParserListener) ExitIps(ctx *IpsContext) {}

// EnterIpv4_address is called when production ipv4_address is entered.
func (s *BaseACIParserListener) EnterIpv4_address(ctx *Ipv4_addressContext) {}

// ExitIpv4_address is called when production ipv4_address is exited.
func (s *BaseACIParserListener) ExitIpv4_address(ctx *Ipv4_addressContext) {}

// EnterIpv6_address is called when production ipv6_address is entered.
func (s *BaseACIParserListener) EnterIpv6_address(ctx *Ipv6_addressContext) {}

// ExitIpv6_address is called when production ipv6_address is exited.
func (s *BaseACIParserListener) ExitIpv6_address(ctx *Ipv6_addressContext) {}

// EnterIpv4 is called when production ipv4 is entered.
func (s *BaseACIParserListener) EnterIpv4(ctx *Ipv4Context) {}

// ExitIpv4 is called when production ipv4 is exited.
func (s *BaseACIParserListener) ExitIpv4(ctx *Ipv4Context) {}

// EnterIpv6 is called when production ipv6 is entered.
func (s *BaseACIParserListener) EnterIpv6(ctx *Ipv6Context) {}

// ExitIpv6 is called when production ipv6 is exited.
func (s *BaseACIParserListener) ExitIpv6(ctx *Ipv6Context) {}

// EnterFqdn is called when production fqdn is entered.
func (s *BaseACIParserListener) EnterFqdn(ctx *FqdnContext) {}

// ExitFqdn is called when production fqdn is exited.
func (s *BaseACIParserListener) ExitFqdn(ctx *FqdnContext) {}

// EnterFullyQualifiedLDAPURI is called when production fullyQualifiedLDAPURI is entered.
func (s *BaseACIParserListener) EnterFullyQualifiedLDAPURI(ctx *FullyQualifiedLDAPURIContext) {}

// ExitFullyQualifiedLDAPURI is called when production fullyQualifiedLDAPURI is exited.
func (s *BaseACIParserListener) ExitFullyQualifiedLDAPURI(ctx *FullyQualifiedLDAPURIContext) {}

// EnterUriSearchFilter is called when production uriSearchFilter is entered.
func (s *BaseACIParserListener) EnterUriSearchFilter(ctx *UriSearchFilterContext) {}

// ExitUriSearchFilter is called when production uriSearchFilter is exited.
func (s *BaseACIParserListener) ExitUriSearchFilter(ctx *UriSearchFilterContext) {}

// EnterUriSearchScopes is called when production uriSearchScopes is entered.
func (s *BaseACIParserListener) EnterUriSearchScopes(ctx *UriSearchScopesContext) {}

// ExitUriSearchScopes is called when production uriSearchScopes is exited.
func (s *BaseACIParserListener) ExitUriSearchScopes(ctx *UriSearchScopesContext) {}

// EnterUriAttributeList is called when production uriAttributeList is entered.
func (s *BaseACIParserListener) EnterUriAttributeList(ctx *UriAttributeListContext) {}

// ExitUriAttributeList is called when production uriAttributeList is exited.
func (s *BaseACIParserListener) ExitUriAttributeList(ctx *UriAttributeListContext) {}

// EnterQuoted_distinguished_name_list is called when production quoted_distinguished_name_list is entered.
func (s *BaseACIParserListener) EnterQuoted_distinguished_name_list(ctx *Quoted_distinguished_name_listContext) {
}

// ExitQuoted_distinguished_name_list is called when production quoted_distinguished_name_list is exited.
func (s *BaseACIParserListener) ExitQuoted_distinguished_name_list(ctx *Quoted_distinguished_name_listContext) {
}

// EnterList_of_quoted_distinguished_names is called when production list_of_quoted_distinguished_names is entered.
func (s *BaseACIParserListener) EnterList_of_quoted_distinguished_names(ctx *List_of_quoted_distinguished_namesContext) {
}

// ExitList_of_quoted_distinguished_names is called when production list_of_quoted_distinguished_names is exited.
func (s *BaseACIParserListener) ExitList_of_quoted_distinguished_names(ctx *List_of_quoted_distinguished_namesContext) {
}

// EnterDn is called when production dn is entered.
func (s *BaseACIParserListener) EnterDn(ctx *DnContext) {}

// ExitDn is called when production dn is exited.
func (s *BaseACIParserListener) ExitDn(ctx *DnContext) {}

// EnterAnonymous_dn_alias is called when production anonymous_dn_alias is entered.
func (s *BaseACIParserListener) EnterAnonymous_dn_alias(ctx *Anonymous_dn_aliasContext) {}

// ExitAnonymous_dn_alias is called when production anonymous_dn_alias is exited.
func (s *BaseACIParserListener) ExitAnonymous_dn_alias(ctx *Anonymous_dn_aliasContext) {}

// EnterAny_user_dn_alias is called when production any_user_dn_alias is entered.
func (s *BaseACIParserListener) EnterAny_user_dn_alias(ctx *Any_user_dn_aliasContext) {}

// ExitAny_user_dn_alias is called when production any_user_dn_alias is exited.
func (s *BaseACIParserListener) ExitAny_user_dn_alias(ctx *Any_user_dn_aliasContext) {}

// EnterSelf_dn_alias is called when production self_dn_alias is entered.
func (s *BaseACIParserListener) EnterSelf_dn_alias(ctx *Self_dn_aliasContext) {}

// ExitSelf_dn_alias is called when production self_dn_alias is exited.
func (s *BaseACIParserListener) ExitSelf_dn_alias(ctx *Self_dn_aliasContext) {}

// EnterParent_dn_alias is called when production parent_dn_alias is entered.
func (s *BaseACIParserListener) EnterParent_dn_alias(ctx *Parent_dn_aliasContext) {}

// ExitParent_dn_alias is called when production parent_dn_alias is exited.
func (s *BaseACIParserListener) ExitParent_dn_alias(ctx *Parent_dn_aliasContext) {}

// EnterRdn_macro is called when production rdn_macro is entered.
func (s *BaseACIParserListener) EnterRdn_macro(ctx *Rdn_macroContext) {}

// ExitRdn_macro is called when production rdn_macro is exited.
func (s *BaseACIParserListener) ExitRdn_macro(ctx *Rdn_macroContext) {}

// EnterParenthetical_filter_expression is called when production parenthetical_filter_expression is entered.
func (s *BaseACIParserListener) EnterParenthetical_filter_expression(ctx *Parenthetical_filter_expressionContext) {
}

// ExitParenthetical_filter_expression is called when production parenthetical_filter_expression is exited.
func (s *BaseACIParserListener) ExitParenthetical_filter_expression(ctx *Parenthetical_filter_expressionContext) {
}

// EnterFilter_expressions is called when production filter_expressions is entered.
func (s *BaseACIParserListener) EnterFilter_expressions(ctx *Filter_expressionsContext) {}

// ExitFilter_expressions is called when production filter_expressions is exited.
func (s *BaseACIParserListener) ExitFilter_expressions(ctx *Filter_expressionsContext) {}

// EnterParenthetical_filter_expression_opt_bool is called when production parenthetical_filter_expression_opt_bool is entered.
func (s *BaseACIParserListener) EnterParenthetical_filter_expression_opt_bool(ctx *Parenthetical_filter_expression_opt_boolContext) {
}

// ExitParenthetical_filter_expression_opt_bool is called when production parenthetical_filter_expression_opt_bool is exited.
func (s *BaseACIParserListener) ExitParenthetical_filter_expression_opt_bool(ctx *Parenthetical_filter_expression_opt_boolContext) {
}

// EnterNot_filter_expression is called when production not_filter_expression is entered.
func (s *BaseACIParserListener) EnterNot_filter_expression(ctx *Not_filter_expressionContext) {}

// ExitNot_filter_expression is called when production not_filter_expression is exited.
func (s *BaseACIParserListener) ExitNot_filter_expression(ctx *Not_filter_expressionContext) {}

// EnterAva_expression is called when production ava_expression is entered.
func (s *BaseACIParserListener) EnterAva_expression(ctx *Ava_expressionContext) {}

// ExitAva_expression is called when production ava_expression is exited.
func (s *BaseACIParserListener) ExitAva_expression(ctx *Ava_expressionContext) {}

// EnterAva_or_rdn is called when production ava_or_rdn is entered.
func (s *BaseACIParserListener) EnterAva_or_rdn(ctx *Ava_or_rdnContext) {}

// ExitAva_or_rdn is called when production ava_or_rdn is exited.
func (s *BaseACIParserListener) ExitAva_or_rdn(ctx *Ava_or_rdnContext) {}

// EnterInheritance_expression is called when production inheritance_expression is entered.
func (s *BaseACIParserListener) EnterInheritance_expression(ctx *Inheritance_expressionContext) {}

// ExitInheritance_expression is called when production inheritance_expression is exited.
func (s *BaseACIParserListener) ExitInheritance_expression(ctx *Inheritance_expressionContext) {}

// EnterInheritance_levels is called when production inheritance_levels is entered.
func (s *BaseACIParserListener) EnterInheritance_levels(ctx *Inheritance_levelsContext) {}

// ExitInheritance_levels is called when production inheritance_levels is exited.
func (s *BaseACIParserListener) ExitInheritance_levels(ctx *Inheritance_levelsContext) {}

// EnterAttr_bind_type_or_value is called when production attr_bind_type_or_value is entered.
func (s *BaseACIParserListener) EnterAttr_bind_type_or_value(ctx *Attr_bind_type_or_valueContext) {}

// ExitAttr_bind_type_or_value is called when production attr_bind_type_or_value is exited.
func (s *BaseACIParserListener) ExitAttr_bind_type_or_value(ctx *Attr_bind_type_or_valueContext) {}

// EnterUSERDN is called when production USERDN is entered.
func (s *BaseACIParserListener) EnterUSERDN(ctx *USERDNContext) {}

// ExitUSERDN is called when production USERDN is exited.
func (s *BaseACIParserListener) ExitUSERDN(ctx *USERDNContext) {}

// EnterGROUPDN is called when production GROUPDN is entered.
func (s *BaseACIParserListener) EnterGROUPDN(ctx *GROUPDNContext) {}

// ExitGROUPDN is called when production GROUPDN is exited.
func (s *BaseACIParserListener) ExitGROUPDN(ctx *GROUPDNContext) {}

// EnterROLEDN is called when production ROLEDN is entered.
func (s *BaseACIParserListener) EnterROLEDN(ctx *ROLEDNContext) {}

// ExitROLEDN is called when production ROLEDN is exited.
func (s *BaseACIParserListener) ExitROLEDN(ctx *ROLEDNContext) {}

// EnterSELFDN is called when production SELFDN is entered.
func (s *BaseACIParserListener) EnterSELFDN(ctx *SELFDNContext) {}

// ExitSELFDN is called when production SELFDN is exited.
func (s *BaseACIParserListener) ExitSELFDN(ctx *SELFDNContext) {}

// EnterLDAPURL is called when production LDAPURL is entered.
func (s *BaseACIParserListener) EnterLDAPURL(ctx *LDAPURLContext) {}

// ExitLDAPURL is called when production LDAPURL is exited.
func (s *BaseACIParserListener) ExitLDAPURL(ctx *LDAPURLContext) {}

// EnterKey_or_value is called when production key_or_value is entered.
func (s *BaseACIParserListener) EnterKey_or_value(ctx *Key_or_valueContext) {}

// ExitKey_or_value is called when production key_or_value is exited.
func (s *BaseACIParserListener) ExitKey_or_value(ctx *Key_or_valueContext) {}

// EnterEqual_to is called when production equal_to is entered.
func (s *BaseACIParserListener) EnterEqual_to(ctx *Equal_toContext) {}

// ExitEqual_to is called when production equal_to is exited.
func (s *BaseACIParserListener) ExitEqual_to(ctx *Equal_toContext) {}

// EnterNot_equal_to is called when production not_equal_to is entered.
func (s *BaseACIParserListener) EnterNot_equal_to(ctx *Not_equal_toContext) {}

// ExitNot_equal_to is called when production not_equal_to is exited.
func (s *BaseACIParserListener) ExitNot_equal_to(ctx *Not_equal_toContext) {}

// EnterGreater_than_or_equal is called when production greater_than_or_equal is entered.
func (s *BaseACIParserListener) EnterGreater_than_or_equal(ctx *Greater_than_or_equalContext) {}

// ExitGreater_than_or_equal is called when production greater_than_or_equal is exited.
func (s *BaseACIParserListener) ExitGreater_than_or_equal(ctx *Greater_than_or_equalContext) {}

// EnterLess_than_or_equal is called when production less_than_or_equal is entered.
func (s *BaseACIParserListener) EnterLess_than_or_equal(ctx *Less_than_or_equalContext) {}

// ExitLess_than_or_equal is called when production less_than_or_equal is exited.
func (s *BaseACIParserListener) ExitLess_than_or_equal(ctx *Less_than_or_equalContext) {}

// EnterApprox is called when production approx is entered.
func (s *BaseACIParserListener) EnterApprox(ctx *ApproxContext) {}

// ExitApprox is called when production approx is exited.
func (s *BaseACIParserListener) ExitApprox(ctx *ApproxContext) {}

// EnterExtensible_rule is called when production extensible_rule is entered.
func (s *BaseACIParserListener) EnterExtensible_rule(ctx *Extensible_ruleContext) {}

// ExitExtensible_rule is called when production extensible_rule is exited.
func (s *BaseACIParserListener) ExitExtensible_rule(ctx *Extensible_ruleContext) {}

// EnterExtensible_rule_with_dn is called when production extensible_rule_with_dn is entered.
func (s *BaseACIParserListener) EnterExtensible_rule_with_dn(ctx *Extensible_rule_with_dnContext) {}

// ExitExtensible_rule_with_dn is called when production extensible_rule_with_dn is exited.
func (s *BaseACIParserListener) ExitExtensible_rule_with_dn(ctx *Extensible_rule_with_dnContext) {}

// EnterExtensible_rule_with_attrs is called when production extensible_rule_with_attrs is entered.
func (s *BaseACIParserListener) EnterExtensible_rule_with_attrs(ctx *Extensible_rule_with_attrsContext) {
}

// ExitExtensible_rule_with_attrs is called when production extensible_rule_with_attrs is exited.
func (s *BaseACIParserListener) ExitExtensible_rule_with_attrs(ctx *Extensible_rule_with_attrsContext) {
}

// EnterExtensible_rule_with_dn_oid is called when production extensible_rule_with_dn_oid is entered.
func (s *BaseACIParserListener) EnterExtensible_rule_with_dn_oid(ctx *Extensible_rule_with_dn_oidContext) {
}

// ExitExtensible_rule_with_dn_oid is called when production extensible_rule_with_dn_oid is exited.
func (s *BaseACIParserListener) ExitExtensible_rule_with_dn_oid(ctx *Extensible_rule_with_dn_oidContext) {
}

// EnterEqualTo is called when production equalTo is entered.
func (s *BaseACIParserListener) EnterEqualTo(ctx *EqualToContext) {}

// ExitEqualTo is called when production equalTo is exited.
func (s *BaseACIParserListener) ExitEqualTo(ctx *EqualToContext) {}

// EnterNotEqualTo is called when production notEqualTo is entered.
func (s *BaseACIParserListener) EnterNotEqualTo(ctx *NotEqualToContext) {}

// ExitNotEqualTo is called when production notEqualTo is exited.
func (s *BaseACIParserListener) ExitNotEqualTo(ctx *NotEqualToContext) {}

// EnterGreaterThan is called when production greaterThan is entered.
func (s *BaseACIParserListener) EnterGreaterThan(ctx *GreaterThanContext) {}

// ExitGreaterThan is called when production greaterThan is exited.
func (s *BaseACIParserListener) ExitGreaterThan(ctx *GreaterThanContext) {}

// EnterLessThan is called when production lessThan is entered.
func (s *BaseACIParserListener) EnterLessThan(ctx *LessThanContext) {}

// ExitLessThan is called when production lessThan is exited.
func (s *BaseACIParserListener) ExitLessThan(ctx *LessThanContext) {}

// EnterGreaterThanOrEqual is called when production greaterThanOrEqual is entered.
func (s *BaseACIParserListener) EnterGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) {}

// ExitGreaterThanOrEqual is called when production greaterThanOrEqual is exited.
func (s *BaseACIParserListener) ExitGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) {}

// EnterLessThanOrEqual is called when production lessThanOrEqual is entered.
func (s *BaseACIParserListener) EnterLessThanOrEqual(ctx *LessThanOrEqualContext) {}

// ExitLessThanOrEqual is called when production lessThanOrEqual is exited.
func (s *BaseACIParserListener) ExitLessThanOrEqual(ctx *LessThanOrEqualContext) {}

// EnterApproximate is called when production approximate is entered.
func (s *BaseACIParserListener) EnterApproximate(ctx *ApproximateContext) {}

// ExitApproximate is called when production approximate is exited.
func (s *BaseACIParserListener) ExitApproximate(ctx *ApproximateContext) {}

// EnterExtensibleRule is called when production extensibleRule is entered.
func (s *BaseACIParserListener) EnterExtensibleRule(ctx *ExtensibleRuleContext) {}

// ExitExtensibleRule is called when production extensibleRule is exited.
func (s *BaseACIParserListener) ExitExtensibleRule(ctx *ExtensibleRuleContext) {}

// EnterExtensibleRuleDNOID is called when production extensibleRuleDNOID is entered.
func (s *BaseACIParserListener) EnterExtensibleRuleDNOID(ctx *ExtensibleRuleDNOIDContext) {}

// ExitExtensibleRuleDNOID is called when production extensibleRuleDNOID is exited.
func (s *BaseACIParserListener) ExitExtensibleRuleDNOID(ctx *ExtensibleRuleDNOIDContext) {}

// EnterExtensibleRuleDN is called when production extensibleRuleDN is entered.
func (s *BaseACIParserListener) EnterExtensibleRuleDN(ctx *ExtensibleRuleDNContext) {}

// ExitExtensibleRuleDN is called when production extensibleRuleDN is exited.
func (s *BaseACIParserListener) ExitExtensibleRuleDN(ctx *ExtensibleRuleDNContext) {}

// EnterExtensibleRuleAttr is called when production extensibleRuleAttr is entered.
func (s *BaseACIParserListener) EnterExtensibleRuleAttr(ctx *ExtensibleRuleAttrContext) {}

// ExitExtensibleRuleAttr is called when production extensibleRuleAttr is exited.
func (s *BaseACIParserListener) ExitExtensibleRuleAttr(ctx *ExtensibleRuleAttrContext) {}

// EnterORDelimiter is called when production oRDelimiter is entered.
func (s *BaseACIParserListener) EnterORDelimiter(ctx *ORDelimiterContext) {}

// ExitORDelimiter is called when production oRDelimiter is exited.
func (s *BaseACIParserListener) ExitORDelimiter(ctx *ORDelimiterContext) {}

// EnterANDDelimiter is called when production aNDDelimiter is entered.
func (s *BaseACIParserListener) EnterANDDelimiter(ctx *ANDDelimiterContext) {}

// ExitANDDelimiter is called when production aNDDelimiter is exited.
func (s *BaseACIParserListener) ExitANDDelimiter(ctx *ANDDelimiterContext) {}
