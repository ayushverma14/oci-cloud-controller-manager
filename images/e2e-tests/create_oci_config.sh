function createOCIConfig() {
    OCI_CONFIG_DIR="$HOME/.oci"
    #
    # args:
    #   1 - oci_key
    #   2 - oci_fingerprint
    #   3 - oci_region
    #   4 - oci_tenancy
    #   5 - oci_user
    #   6 - oci profile name
    # Create config directory.
    mkdir -p ${OCI_CONFIG_DIR}
    if [ $? -ne 0 ]; then
         echo "Could not create oci config directory at ${OCI_CONFIG_DIR}"
         exit 1
    fi
    echo "Created OCI config directory at ${OCI_CONFIG_DIR}"

    # Create OCI key (PEM) file.
    KEY_PEM_FILE="${OCI_CONFIG_DIR}/${6}_oci_api_key.pem"

    echo $1 | sed 's/ //g' | openssl enc -base64 -d -A > $KEY_PEM_FILE
    echo "Created oci key file at $KEY_PEM_FILE"
    oci setup repair-file-permissions --file ${KEY_PEM_FILE}
    # Create OCI config file.
    CONFIG_FILE="${OCI_CONFIG_DIR}/config"
    CONFIG_CONTENT="[$6]\nuser=$5\nfingerprint=$2\nkey_file=$KEY_PEM_FILE\ntenancy=$4\nregion=$3\n"
    echo -e $CONFIG_CONTENT >> $CONFIG_FILE
    echo "Created oci config file at $CONFIG_FILE"
    export export CONFIG_FILE
    oci setup repair-file-permissions --file ${CONFIG_FILE}

}


createOCIConfig $OCI_KEY $OCI_FINGERPRINT $OCI_REGION $OCI_TENANCY $OCI_USER DEFAULT
createOCIConfig $SOPS_OCI_KEY $SOPS_OCI_FINGERPRINT $SOPS_OCI_REGION $SOPS_OCI_TENANCY $SOPS_OCI_USER SOPS
