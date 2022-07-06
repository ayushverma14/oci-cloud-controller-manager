upload_logs() {
  exit_status=$?
  touch /var/log/ccm-e2e-logs.txt
  oci os object put -ns odx-mockcustomer -bn oke-dp-e2e-logs --name "${ODO_APPLICATION_CHANGE}_ccm_e2e_logs.txt" --file /var/log/ccm-e2e-logs.txt
  exit $exit_status
}
# Upload logs to object storage on exit
trap upload_logs EXIT


set -o pipefail
# SOPS is used to decrypt secrets
/usr/local/bin/create_oci_config.sh &&
sops -i -d --oci-kms ${KMS_SECRETS_KEY} --oci-profile SOPS temp_repos/secrets.tar.gz &&
mkdir -p temp_repos/secrets &&
tar -xzf temp_repos/secrets.tar.gz -C temp_repos/secrets &&
mv temp_repos/secrets / &&
${OKE_TEST_COMMAND} | tee -a /var/log/ccm-e2e-logs.txt
