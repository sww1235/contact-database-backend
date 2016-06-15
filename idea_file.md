use 'object oriented' linking for contact cards

### heirarchy
-   Database
    -   groups
        -   family
            -   individual
      
individuals inherit any info from family such as address unless specified in indivdual

families do not inherit from groups

have a default "family" with no inheritable info for individuals with no family associated

use vCard format for individuals, potentially extended with custom attributes for family/org. Each individual will be a separate file that stores family relationships which are then parsed at runtime.
