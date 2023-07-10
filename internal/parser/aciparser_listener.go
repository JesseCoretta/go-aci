// Code generated from ACIParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ACIParser
import "github.com/antlr4-go/antlr/v4"

// ACIParserListener is a complete listener for a parse tree produced by ACIParser.
type ACIParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterAccessControlInstruction is called when entering the accessControlInstruction production.
	EnterAccessControlInstruction(c *AccessControlInstructionContext)

	// EnterAccessControlInstructionNameAndVersion is called when entering the accessControlInstructionNameAndVersion production.
	EnterAccessControlInstructionNameAndVersion(c *AccessControlInstructionNameAndVersionContext)

	// EnterPermissionBindRules is called when entering the permissionBindRules production.
	EnterPermissionBindRules(c *PermissionBindRulesContext)

	// EnterPermissionBindRule is called when entering the permissionBindRule production.
	EnterPermissionBindRule(c *PermissionBindRuleContext)

	// EnterPermissionExpression is called when entering the permissionExpression production.
	EnterPermissionExpression(c *PermissionExpressionContext)

	// EnterTargetRuleExpressions is called when entering the targetRuleExpressions production.
	EnterTargetRuleExpressions(c *TargetRuleExpressionsContext)

	// EnterTargetcontrol is called when entering the targetcontrol production.
	EnterTargetcontrol(c *TargetcontrolContext)

	// EnterExtop is called when entering the extop production.
	EnterExtop(c *ExtopContext)

	// EnterTargetfilter is called when entering the targetfilter production.
	EnterTargetfilter(c *TargetfilterContext)

	// EnterTargattrfilters is called when entering the targattrfilters production.
	EnterTargattrfilters(c *TargattrfiltersContext)

	// EnterTargetscope is called when entering the targetscope production.
	EnterTargetscope(c *TargetscopeContext)

	// EnterTargetattr is called when entering the targetattr production.
	EnterTargetattr(c *TargetattrContext)

	// EnterTargetdn is called when entering the targetdn production.
	EnterTargetdn(c *TargetdnContext)

	// EnterTargettodn is called when entering the targettodn production.
	EnterTargettodn(c *TargettodnContext)

	// EnterTargetfromDN is called when entering the targetfromDN production.
	EnterTargetfromDN(c *TargetfromDNContext)

	// EnterParentheticalControls is called when entering the parentheticalControls production.
	EnterParentheticalControls(c *ParentheticalControlsContext)

	// EnterParentheticalExtendedOperations is called when entering the parentheticalExtendedOperations production.
	EnterParentheticalExtendedOperations(c *ParentheticalExtendedOperationsContext)

	// EnterParentheticalTargetFilterExpression is called when entering the parentheticalTargetFilterExpression production.
	EnterParentheticalTargetFilterExpression(c *ParentheticalTargetFilterExpressionContext)

	// EnterQuotedFilterExpression is called when entering the quotedFilterExpression production.
	EnterQuotedFilterExpression(c *QuotedFilterExpressionContext)

	// EnterTargetDistinguishedNames is called when entering the targetDistinguishedNames production.
	EnterTargetDistinguishedNames(c *TargetDistinguishedNamesContext)

	// EnterTargetToDistinguishedName is called when entering the targetToDistinguishedName production.
	EnterTargetToDistinguishedName(c *TargetToDistinguishedNameContext)

	// EnterTargetFromDistinguishedName is called when entering the targetFromDistinguishedName production.
	EnterTargetFromDistinguishedName(c *TargetFromDistinguishedNameContext)

	// EnterParentheticalTargetAttrFilters is called when entering the parentheticalTargetAttrFilters production.
	EnterParentheticalTargetAttrFilters(c *ParentheticalTargetAttrFiltersContext)

	// EnterQuotedAttributeFilters is called when entering the quotedAttributeFilters production.
	EnterQuotedAttributeFilters(c *QuotedAttributeFiltersContext)

	// EnterQuotedAttributeFilterSet is called when entering the quotedAttributeFilterSet production.
	EnterQuotedAttributeFilterSet(c *QuotedAttributeFilterSetContext)

	// EnterQuotedAttributeFilter is called when entering the quotedAttributeFilter production.
	EnterQuotedAttributeFilter(c *QuotedAttributeFilterContext)

	// EnterTargetScopeBindRule is called when entering the targetScopeBindRule production.
	EnterTargetScopeBindRule(c *TargetScopeBindRuleContext)

	// EnterTargetAttrBindRule is called when entering the targetAttrBindRule production.
	EnterTargetAttrBindRule(c *TargetAttrBindRuleContext)

	// EnterAttributeTypesList is called when entering the attributeTypesList production.
	EnterAttributeTypesList(c *AttributeTypesListContext)

	// EnterBindRuleInstance is called when entering the bindRuleInstance production.
	EnterBindRuleInstance(c *BindRuleInstanceContext)

	// EnterParentheticalBindRuleInstanceWithRequiredBooleanOperator is called when entering the parentheticalBindRuleInstanceWithRequiredBooleanOperator production.
	EnterParentheticalBindRuleInstanceWithRequiredBooleanOperator(c *ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext)

	// EnterParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursion is called when entering the parentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursion production.
	EnterParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursion(c *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext)

	// EnterNegatedBindRuleExpressionRecursion is called when entering the negatedBindRuleExpressionRecursion production.
	EnterNegatedBindRuleExpressionRecursion(c *NegatedBindRuleExpressionRecursionContext)

	// EnterParentheticalBindRuleExpressionRecursion is called when entering the parentheticalBindRuleExpressionRecursion production.
	EnterParentheticalBindRuleExpressionRecursion(c *ParentheticalBindRuleExpressionRecursionContext)

	// EnterBindRuleExpressionRecursion is called when entering the bindRuleExpressionRecursion production.
	EnterBindRuleExpressionRecursion(c *BindRuleExpressionRecursionContext)

	// EnterBindRuleExpression is called when entering the bindRuleExpression production.
	EnterBindRuleExpression(c *BindRuleExpressionContext)

	// EnterUserdn is called when entering the userdn production.
	EnterUserdn(c *UserdnContext)

	// EnterUserattr is called when entering the userattr production.
	EnterUserattr(c *UserattrContext)

	// EnterGroupdn is called when entering the groupdn production.
	EnterGroupdn(c *GroupdnContext)

	// EnterGroupattr is called when entering the groupattr production.
	EnterGroupattr(c *GroupattrContext)

	// EnterRoledn is called when entering the roledn production.
	EnterRoledn(c *RolednContext)

	// EnterDns is called when entering the dns production.
	EnterDns(c *DnsContext)

	// EnterIp is called when entering the ip production.
	EnterIp(c *IpContext)

	// EnterTimeofday is called when entering the timeofday production.
	EnterTimeofday(c *TimeofdayContext)

	// EnterDayofweek is called when entering the dayofweek production.
	EnterDayofweek(c *DayofweekContext)

	// EnterSsf is called when entering the ssf production.
	EnterSsf(c *SsfContext)

	// EnterAuthmethod is called when entering the authmethod production.
	EnterAuthmethod(c *AuthmethodContext)

	// EnterParentheticalUserDistinguishedName is called when entering the parentheticalUserDistinguishedName production.
	EnterParentheticalUserDistinguishedName(c *ParentheticalUserDistinguishedNameContext)

	// EnterUserDistinguishedNameExpression is called when entering the userDistinguishedNameExpression production.
	EnterUserDistinguishedNameExpression(c *UserDistinguishedNameExpressionContext)

	// EnterParentheticalRoleDistinguishedName is called when entering the parentheticalRoleDistinguishedName production.
	EnterParentheticalRoleDistinguishedName(c *ParentheticalRoleDistinguishedNameContext)

	// EnterRoleDistinguishedNameExpression is called when entering the roleDistinguishedNameExpression production.
	EnterRoleDistinguishedNameExpression(c *RoleDistinguishedNameExpressionContext)

	// EnterParentheticalGroupDistinguishedName is called when entering the parentheticalGroupDistinguishedName production.
	EnterParentheticalGroupDistinguishedName(c *ParentheticalGroupDistinguishedNameContext)

	// EnterGroupDistinguishedNameExpression is called when entering the groupDistinguishedNameExpression production.
	EnterGroupDistinguishedNameExpression(c *GroupDistinguishedNameExpressionContext)

	// EnterParentheticalUserAttributes is called when entering the parentheticalUserAttributes production.
	EnterParentheticalUserAttributes(c *ParentheticalUserAttributesContext)

	// EnterUserAttributesExpression is called when entering the userAttributesExpression production.
	EnterUserAttributesExpression(c *UserAttributesExpressionContext)

	// EnterParentheticalGroupAttributes is called when entering the parentheticalGroupAttributes production.
	EnterParentheticalGroupAttributes(c *ParentheticalGroupAttributesContext)

	// EnterGroupAttributesExpression is called when entering the groupAttributesExpression production.
	EnterGroupAttributesExpression(c *GroupAttributesExpressionContext)

	// EnterParentheticalAuthenticationMethod is called when entering the parentheticalAuthenticationMethod production.
	EnterParentheticalAuthenticationMethod(c *ParentheticalAuthenticationMethodContext)

	// EnterAuthenticationMethodExpression is called when entering the authenticationMethodExpression production.
	EnterAuthenticationMethodExpression(c *AuthenticationMethodExpressionContext)

	// EnterParentheticalDNS is called when entering the parentheticalDNS production.
	EnterParentheticalDNS(c *ParentheticalDNSContext)

	// EnterDNSBindRule is called when entering the dNSBindRule production.
	EnterDNSBindRule(c *DNSBindRuleContext)

	// EnterParentheticalTimeOfDay is called when entering the parentheticalTimeOfDay production.
	EnterParentheticalTimeOfDay(c *ParentheticalTimeOfDayContext)

	// EnterTimeOfDayBindRule is called when entering the timeOfDayBindRule production.
	EnterTimeOfDayBindRule(c *TimeOfDayBindRuleContext)

	// EnterParentheticalDayOfWeek is called when entering the parentheticalDayOfWeek production.
	EnterParentheticalDayOfWeek(c *ParentheticalDayOfWeekContext)

	// EnterDayOfWeekExpression is called when entering the dayOfWeekExpression production.
	EnterDayOfWeekExpression(c *DayOfWeekExpressionContext)

	// EnterParentheticalIPAddress is called when entering the parentheticalIPAddress production.
	EnterParentheticalIPAddress(c *ParentheticalIPAddressContext)

	// EnterIpAddressBindRule is called when entering the ipAddressBindRule production.
	EnterIpAddressBindRule(c *IpAddressBindRuleContext)

	// EnterParentheticalSecurityStrengthFactor is called when entering the parentheticalSecurityStrengthFactor production.
	EnterParentheticalSecurityStrengthFactor(c *ParentheticalSecurityStrengthFactorContext)

	// EnterSecurityStrengthFactorExpression is called when entering the securityStrengthFactorExpression production.
	EnterSecurityStrengthFactorExpression(c *SecurityStrengthFactorExpressionContext)

	// EnterDayOfWeekValue is called when entering the dayOfWeekValue production.
	EnterDayOfWeekValue(c *DayOfWeekValueContext)

	// EnterFullyQualifiedDomainNameValue is called when entering the fullyQualifiedDomainNameValue production.
	EnterFullyQualifiedDomainNameValue(c *FullyQualifiedDomainNameValueContext)

	// EnterObjectIdentifierValues is called when entering the objectIdentifierValues production.
	EnterObjectIdentifierValues(c *ObjectIdentifierValuesContext)

	// EnterObjectIdentifierValue is called when entering the objectIdentifierValue production.
	EnterObjectIdentifierValue(c *ObjectIdentifierValueContext)

	// EnterIPV6AddressValue is called when entering the iPV6AddressValue production.
	EnterIPV6AddressValue(c *IPV6AddressValueContext)

	// EnterIPV4AddressValue is called when entering the iPV4AddressValue production.
	EnterIPV4AddressValue(c *IPV4AddressValueContext)

	// EnterSecurityStrengthFactorValue is called when entering the securityStrengthFactorValue production.
	EnterSecurityStrengthFactorValue(c *SecurityStrengthFactorValueContext)

	// EnterTimeOfDayValue is called when entering the timeOfDayValue production.
	EnterTimeOfDayValue(c *TimeOfDayValueContext)

	// EnterObjectIdentifierArc is called when entering the objectIdentifierArc production.
	EnterObjectIdentifierArc(c *ObjectIdentifierArcContext)

	// EnterInheritanceExpression is called when entering the inheritanceExpression production.
	EnterInheritanceExpression(c *InheritanceExpressionContext)

	// EnterInheritanceLevelValue is called when entering the inheritanceLevelValue production.
	EnterInheritanceLevelValue(c *InheritanceLevelValueContext)

	// EnterAttributeBindTypeOrValueValue is called when entering the attributeBindTypeOrValueValue production.
	EnterAttributeBindTypeOrValueValue(c *AttributeBindTypeOrValueValueContext)

	// EnterAttributeFiltersExpression is called when entering the attributeFiltersExpression production.
	EnterAttributeFiltersExpression(c *AttributeFiltersExpressionContext)

	// EnterAttributeFilterSetExpression is called when entering the attributeFilterSetExpression production.
	EnterAttributeFilterSetExpression(c *AttributeFilterSetExpressionContext)

	// EnterDoubleAmpersandDelimiter is called when entering the doubleAmpersandDelimiter production.
	EnterDoubleAmpersandDelimiter(c *DoubleAmpersandDelimiterContext)

	// EnterAttributeFilterExpression is called when entering the attributeFilterExpression production.
	EnterAttributeFilterExpression(c *AttributeFilterExpressionContext)

	// EnterDistinguishedNamesList is called when entering the distinguishedNamesList production.
	EnterDistinguishedNamesList(c *DistinguishedNamesListContext)

	// EnterDoublePipeDelimiter is called when entering the doublePipeDelimiter production.
	EnterDoublePipeDelimiter(c *DoublePipeDelimiterContext)

	// EnterUriAndBindType is called when entering the uriAndBindType production.
	EnterUriAndBindType(c *UriAndBindTypeContext)

	// EnterFullyQualifiedLDAPURI is called when entering the fullyQualifiedLDAPURI production.
	EnterFullyQualifiedLDAPURI(c *FullyQualifiedLDAPURIContext)

	// EnterUriSearchFilter is called when entering the uriSearchFilter production.
	EnterUriSearchFilter(c *UriSearchFilterContext)

	// EnterUriSearchScopes is called when entering the uriSearchScopes production.
	EnterUriSearchScopes(c *UriSearchScopesContext)

	// EnterUriAttributeList is called when entering the uriAttributeList production.
	EnterUriAttributeList(c *UriAttributeListContext)

	// EnterDistinguishedNameValue is called when entering the distinguishedNameValue production.
	EnterDistinguishedNameValue(c *DistinguishedNameValueContext)

	// EnterRelativeDistinguishedNameValue is called when entering the relativeDistinguishedNameValue production.
	EnterRelativeDistinguishedNameValue(c *RelativeDistinguishedNameValueContext)

	// EnterRelativeDistinguishedNameMacro is called when entering the relativeDistinguishedNameMacro production.
	EnterRelativeDistinguishedNameMacro(c *RelativeDistinguishedNameMacroContext)

	// EnterParentheticalFilterExpression is called when entering the parentheticalFilterExpression production.
	EnterParentheticalFilterExpression(c *ParentheticalFilterExpressionContext)

	// EnterFilterExpressions is called when entering the filterExpressions production.
	EnterFilterExpressions(c *FilterExpressionsContext)

	// EnterParentheticalFilterExpressionWithOptionalBooleanOperator is called when entering the parentheticalFilterExpressionWithOptionalBooleanOperator production.
	EnterParentheticalFilterExpressionWithOptionalBooleanOperator(c *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext)

	// EnterNegatedFilterExpression is called when entering the negatedFilterExpression production.
	EnterNegatedFilterExpression(c *NegatedFilterExpressionContext)

	// EnterAttributeValueAssertionExpression is called when entering the attributeValueAssertionExpression production.
	EnterAttributeValueAssertionExpression(c *AttributeValueAssertionExpressionContext)

	// EnterAttributeValueAssertionStatement is called when entering the attributeValueAssertionStatement production.
	EnterAttributeValueAssertionStatement(c *AttributeValueAssertionStatementContext)

	// EnterAttributeTypeIdentifier is called when entering the attributeTypeIdentifier production.
	EnterAttributeTypeIdentifier(c *AttributeTypeIdentifierContext)

	// EnterAttributeAssertionValue is called when entering the attributeAssertionValue production.
	EnterAttributeAssertionValue(c *AttributeAssertionValueContext)

	// EnterAttributeOperators is called when entering the attributeOperators production.
	EnterAttributeOperators(c *AttributeOperatorsContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitAccessControlInstruction is called when exiting the accessControlInstruction production.
	ExitAccessControlInstruction(c *AccessControlInstructionContext)

	// ExitAccessControlInstructionNameAndVersion is called when exiting the accessControlInstructionNameAndVersion production.
	ExitAccessControlInstructionNameAndVersion(c *AccessControlInstructionNameAndVersionContext)

	// ExitPermissionBindRules is called when exiting the permissionBindRules production.
	ExitPermissionBindRules(c *PermissionBindRulesContext)

	// ExitPermissionBindRule is called when exiting the permissionBindRule production.
	ExitPermissionBindRule(c *PermissionBindRuleContext)

	// ExitPermissionExpression is called when exiting the permissionExpression production.
	ExitPermissionExpression(c *PermissionExpressionContext)

	// ExitTargetRuleExpressions is called when exiting the targetRuleExpressions production.
	ExitTargetRuleExpressions(c *TargetRuleExpressionsContext)

	// ExitTargetcontrol is called when exiting the targetcontrol production.
	ExitTargetcontrol(c *TargetcontrolContext)

	// ExitExtop is called when exiting the extop production.
	ExitExtop(c *ExtopContext)

	// ExitTargetfilter is called when exiting the targetfilter production.
	ExitTargetfilter(c *TargetfilterContext)

	// ExitTargattrfilters is called when exiting the targattrfilters production.
	ExitTargattrfilters(c *TargattrfiltersContext)

	// ExitTargetscope is called when exiting the targetscope production.
	ExitTargetscope(c *TargetscopeContext)

	// ExitTargetattr is called when exiting the targetattr production.
	ExitTargetattr(c *TargetattrContext)

	// ExitTargetdn is called when exiting the targetdn production.
	ExitTargetdn(c *TargetdnContext)

	// ExitTargettodn is called when exiting the targettodn production.
	ExitTargettodn(c *TargettodnContext)

	// ExitTargetfromDN is called when exiting the targetfromDN production.
	ExitTargetfromDN(c *TargetfromDNContext)

	// ExitParentheticalControls is called when exiting the parentheticalControls production.
	ExitParentheticalControls(c *ParentheticalControlsContext)

	// ExitParentheticalExtendedOperations is called when exiting the parentheticalExtendedOperations production.
	ExitParentheticalExtendedOperations(c *ParentheticalExtendedOperationsContext)

	// ExitParentheticalTargetFilterExpression is called when exiting the parentheticalTargetFilterExpression production.
	ExitParentheticalTargetFilterExpression(c *ParentheticalTargetFilterExpressionContext)

	// ExitQuotedFilterExpression is called when exiting the quotedFilterExpression production.
	ExitQuotedFilterExpression(c *QuotedFilterExpressionContext)

	// ExitTargetDistinguishedNames is called when exiting the targetDistinguishedNames production.
	ExitTargetDistinguishedNames(c *TargetDistinguishedNamesContext)

	// ExitTargetToDistinguishedName is called when exiting the targetToDistinguishedName production.
	ExitTargetToDistinguishedName(c *TargetToDistinguishedNameContext)

	// ExitTargetFromDistinguishedName is called when exiting the targetFromDistinguishedName production.
	ExitTargetFromDistinguishedName(c *TargetFromDistinguishedNameContext)

	// ExitParentheticalTargetAttrFilters is called when exiting the parentheticalTargetAttrFilters production.
	ExitParentheticalTargetAttrFilters(c *ParentheticalTargetAttrFiltersContext)

	// ExitQuotedAttributeFilters is called when exiting the quotedAttributeFilters production.
	ExitQuotedAttributeFilters(c *QuotedAttributeFiltersContext)

	// ExitQuotedAttributeFilterSet is called when exiting the quotedAttributeFilterSet production.
	ExitQuotedAttributeFilterSet(c *QuotedAttributeFilterSetContext)

	// ExitQuotedAttributeFilter is called when exiting the quotedAttributeFilter production.
	ExitQuotedAttributeFilter(c *QuotedAttributeFilterContext)

	// ExitTargetScopeBindRule is called when exiting the targetScopeBindRule production.
	ExitTargetScopeBindRule(c *TargetScopeBindRuleContext)

	// ExitTargetAttrBindRule is called when exiting the targetAttrBindRule production.
	ExitTargetAttrBindRule(c *TargetAttrBindRuleContext)

	// ExitAttributeTypesList is called when exiting the attributeTypesList production.
	ExitAttributeTypesList(c *AttributeTypesListContext)

	// ExitBindRuleInstance is called when exiting the bindRuleInstance production.
	ExitBindRuleInstance(c *BindRuleInstanceContext)

	// ExitParentheticalBindRuleInstanceWithRequiredBooleanOperator is called when exiting the parentheticalBindRuleInstanceWithRequiredBooleanOperator production.
	ExitParentheticalBindRuleInstanceWithRequiredBooleanOperator(c *ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext)

	// ExitParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursion is called when exiting the parentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursion production.
	ExitParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursion(c *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext)

	// ExitNegatedBindRuleExpressionRecursion is called when exiting the negatedBindRuleExpressionRecursion production.
	ExitNegatedBindRuleExpressionRecursion(c *NegatedBindRuleExpressionRecursionContext)

	// ExitParentheticalBindRuleExpressionRecursion is called when exiting the parentheticalBindRuleExpressionRecursion production.
	ExitParentheticalBindRuleExpressionRecursion(c *ParentheticalBindRuleExpressionRecursionContext)

	// ExitBindRuleExpressionRecursion is called when exiting the bindRuleExpressionRecursion production.
	ExitBindRuleExpressionRecursion(c *BindRuleExpressionRecursionContext)

	// ExitBindRuleExpression is called when exiting the bindRuleExpression production.
	ExitBindRuleExpression(c *BindRuleExpressionContext)

	// ExitUserdn is called when exiting the userdn production.
	ExitUserdn(c *UserdnContext)

	// ExitUserattr is called when exiting the userattr production.
	ExitUserattr(c *UserattrContext)

	// ExitGroupdn is called when exiting the groupdn production.
	ExitGroupdn(c *GroupdnContext)

	// ExitGroupattr is called when exiting the groupattr production.
	ExitGroupattr(c *GroupattrContext)

	// ExitRoledn is called when exiting the roledn production.
	ExitRoledn(c *RolednContext)

	// ExitDns is called when exiting the dns production.
	ExitDns(c *DnsContext)

	// ExitIp is called when exiting the ip production.
	ExitIp(c *IpContext)

	// ExitTimeofday is called when exiting the timeofday production.
	ExitTimeofday(c *TimeofdayContext)

	// ExitDayofweek is called when exiting the dayofweek production.
	ExitDayofweek(c *DayofweekContext)

	// ExitSsf is called when exiting the ssf production.
	ExitSsf(c *SsfContext)

	// ExitAuthmethod is called when exiting the authmethod production.
	ExitAuthmethod(c *AuthmethodContext)

	// ExitParentheticalUserDistinguishedName is called when exiting the parentheticalUserDistinguishedName production.
	ExitParentheticalUserDistinguishedName(c *ParentheticalUserDistinguishedNameContext)

	// ExitUserDistinguishedNameExpression is called when exiting the userDistinguishedNameExpression production.
	ExitUserDistinguishedNameExpression(c *UserDistinguishedNameExpressionContext)

	// ExitParentheticalRoleDistinguishedName is called when exiting the parentheticalRoleDistinguishedName production.
	ExitParentheticalRoleDistinguishedName(c *ParentheticalRoleDistinguishedNameContext)

	// ExitRoleDistinguishedNameExpression is called when exiting the roleDistinguishedNameExpression production.
	ExitRoleDistinguishedNameExpression(c *RoleDistinguishedNameExpressionContext)

	// ExitParentheticalGroupDistinguishedName is called when exiting the parentheticalGroupDistinguishedName production.
	ExitParentheticalGroupDistinguishedName(c *ParentheticalGroupDistinguishedNameContext)

	// ExitGroupDistinguishedNameExpression is called when exiting the groupDistinguishedNameExpression production.
	ExitGroupDistinguishedNameExpression(c *GroupDistinguishedNameExpressionContext)

	// ExitParentheticalUserAttributes is called when exiting the parentheticalUserAttributes production.
	ExitParentheticalUserAttributes(c *ParentheticalUserAttributesContext)

	// ExitUserAttributesExpression is called when exiting the userAttributesExpression production.
	ExitUserAttributesExpression(c *UserAttributesExpressionContext)

	// ExitParentheticalGroupAttributes is called when exiting the parentheticalGroupAttributes production.
	ExitParentheticalGroupAttributes(c *ParentheticalGroupAttributesContext)

	// ExitGroupAttributesExpression is called when exiting the groupAttributesExpression production.
	ExitGroupAttributesExpression(c *GroupAttributesExpressionContext)

	// ExitParentheticalAuthenticationMethod is called when exiting the parentheticalAuthenticationMethod production.
	ExitParentheticalAuthenticationMethod(c *ParentheticalAuthenticationMethodContext)

	// ExitAuthenticationMethodExpression is called when exiting the authenticationMethodExpression production.
	ExitAuthenticationMethodExpression(c *AuthenticationMethodExpressionContext)

	// ExitParentheticalDNS is called when exiting the parentheticalDNS production.
	ExitParentheticalDNS(c *ParentheticalDNSContext)

	// ExitDNSBindRule is called when exiting the dNSBindRule production.
	ExitDNSBindRule(c *DNSBindRuleContext)

	// ExitParentheticalTimeOfDay is called when exiting the parentheticalTimeOfDay production.
	ExitParentheticalTimeOfDay(c *ParentheticalTimeOfDayContext)

	// ExitTimeOfDayBindRule is called when exiting the timeOfDayBindRule production.
	ExitTimeOfDayBindRule(c *TimeOfDayBindRuleContext)

	// ExitParentheticalDayOfWeek is called when exiting the parentheticalDayOfWeek production.
	ExitParentheticalDayOfWeek(c *ParentheticalDayOfWeekContext)

	// ExitDayOfWeekExpression is called when exiting the dayOfWeekExpression production.
	ExitDayOfWeekExpression(c *DayOfWeekExpressionContext)

	// ExitParentheticalIPAddress is called when exiting the parentheticalIPAddress production.
	ExitParentheticalIPAddress(c *ParentheticalIPAddressContext)

	// ExitIpAddressBindRule is called when exiting the ipAddressBindRule production.
	ExitIpAddressBindRule(c *IpAddressBindRuleContext)

	// ExitParentheticalSecurityStrengthFactor is called when exiting the parentheticalSecurityStrengthFactor production.
	ExitParentheticalSecurityStrengthFactor(c *ParentheticalSecurityStrengthFactorContext)

	// ExitSecurityStrengthFactorExpression is called when exiting the securityStrengthFactorExpression production.
	ExitSecurityStrengthFactorExpression(c *SecurityStrengthFactorExpressionContext)

	// ExitDayOfWeekValue is called when exiting the dayOfWeekValue production.
	ExitDayOfWeekValue(c *DayOfWeekValueContext)

	// ExitFullyQualifiedDomainNameValue is called when exiting the fullyQualifiedDomainNameValue production.
	ExitFullyQualifiedDomainNameValue(c *FullyQualifiedDomainNameValueContext)

	// ExitObjectIdentifierValues is called when exiting the objectIdentifierValues production.
	ExitObjectIdentifierValues(c *ObjectIdentifierValuesContext)

	// ExitObjectIdentifierValue is called when exiting the objectIdentifierValue production.
	ExitObjectIdentifierValue(c *ObjectIdentifierValueContext)

	// ExitIPV6AddressValue is called when exiting the iPV6AddressValue production.
	ExitIPV6AddressValue(c *IPV6AddressValueContext)

	// ExitIPV4AddressValue is called when exiting the iPV4AddressValue production.
	ExitIPV4AddressValue(c *IPV4AddressValueContext)

	// ExitSecurityStrengthFactorValue is called when exiting the securityStrengthFactorValue production.
	ExitSecurityStrengthFactorValue(c *SecurityStrengthFactorValueContext)

	// ExitTimeOfDayValue is called when exiting the timeOfDayValue production.
	ExitTimeOfDayValue(c *TimeOfDayValueContext)

	// ExitObjectIdentifierArc is called when exiting the objectIdentifierArc production.
	ExitObjectIdentifierArc(c *ObjectIdentifierArcContext)

	// ExitInheritanceExpression is called when exiting the inheritanceExpression production.
	ExitInheritanceExpression(c *InheritanceExpressionContext)

	// ExitInheritanceLevelValue is called when exiting the inheritanceLevelValue production.
	ExitInheritanceLevelValue(c *InheritanceLevelValueContext)

	// ExitAttributeBindTypeOrValueValue is called when exiting the attributeBindTypeOrValueValue production.
	ExitAttributeBindTypeOrValueValue(c *AttributeBindTypeOrValueValueContext)

	// ExitAttributeFiltersExpression is called when exiting the attributeFiltersExpression production.
	ExitAttributeFiltersExpression(c *AttributeFiltersExpressionContext)

	// ExitAttributeFilterSetExpression is called when exiting the attributeFilterSetExpression production.
	ExitAttributeFilterSetExpression(c *AttributeFilterSetExpressionContext)

	// ExitDoubleAmpersandDelimiter is called when exiting the doubleAmpersandDelimiter production.
	ExitDoubleAmpersandDelimiter(c *DoubleAmpersandDelimiterContext)

	// ExitAttributeFilterExpression is called when exiting the attributeFilterExpression production.
	ExitAttributeFilterExpression(c *AttributeFilterExpressionContext)

	// ExitDistinguishedNamesList is called when exiting the distinguishedNamesList production.
	ExitDistinguishedNamesList(c *DistinguishedNamesListContext)

	// ExitDoublePipeDelimiter is called when exiting the doublePipeDelimiter production.
	ExitDoublePipeDelimiter(c *DoublePipeDelimiterContext)

	// ExitUriAndBindType is called when exiting the uriAndBindType production.
	ExitUriAndBindType(c *UriAndBindTypeContext)

	// ExitFullyQualifiedLDAPURI is called when exiting the fullyQualifiedLDAPURI production.
	ExitFullyQualifiedLDAPURI(c *FullyQualifiedLDAPURIContext)

	// ExitUriSearchFilter is called when exiting the uriSearchFilter production.
	ExitUriSearchFilter(c *UriSearchFilterContext)

	// ExitUriSearchScopes is called when exiting the uriSearchScopes production.
	ExitUriSearchScopes(c *UriSearchScopesContext)

	// ExitUriAttributeList is called when exiting the uriAttributeList production.
	ExitUriAttributeList(c *UriAttributeListContext)

	// ExitDistinguishedNameValue is called when exiting the distinguishedNameValue production.
	ExitDistinguishedNameValue(c *DistinguishedNameValueContext)

	// ExitRelativeDistinguishedNameValue is called when exiting the relativeDistinguishedNameValue production.
	ExitRelativeDistinguishedNameValue(c *RelativeDistinguishedNameValueContext)

	// ExitRelativeDistinguishedNameMacro is called when exiting the relativeDistinguishedNameMacro production.
	ExitRelativeDistinguishedNameMacro(c *RelativeDistinguishedNameMacroContext)

	// ExitParentheticalFilterExpression is called when exiting the parentheticalFilterExpression production.
	ExitParentheticalFilterExpression(c *ParentheticalFilterExpressionContext)

	// ExitFilterExpressions is called when exiting the filterExpressions production.
	ExitFilterExpressions(c *FilterExpressionsContext)

	// ExitParentheticalFilterExpressionWithOptionalBooleanOperator is called when exiting the parentheticalFilterExpressionWithOptionalBooleanOperator production.
	ExitParentheticalFilterExpressionWithOptionalBooleanOperator(c *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext)

	// ExitNegatedFilterExpression is called when exiting the negatedFilterExpression production.
	ExitNegatedFilterExpression(c *NegatedFilterExpressionContext)

	// ExitAttributeValueAssertionExpression is called when exiting the attributeValueAssertionExpression production.
	ExitAttributeValueAssertionExpression(c *AttributeValueAssertionExpressionContext)

	// ExitAttributeValueAssertionStatement is called when exiting the attributeValueAssertionStatement production.
	ExitAttributeValueAssertionStatement(c *AttributeValueAssertionStatementContext)

	// ExitAttributeTypeIdentifier is called when exiting the attributeTypeIdentifier production.
	ExitAttributeTypeIdentifier(c *AttributeTypeIdentifierContext)

	// ExitAttributeAssertionValue is called when exiting the attributeAssertionValue production.
	ExitAttributeAssertionValue(c *AttributeAssertionValueContext)

	// ExitAttributeOperators is called when exiting the attributeOperators production.
	ExitAttributeOperators(c *AttributeOperatorsContext)
}
