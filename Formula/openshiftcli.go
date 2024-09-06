package main

// OpenshiftCli - OpenShift command-line interface tools
// Homepage: https://www.openshift.com/

import (
	"fmt"
	
	"os/exec"
)

func installOpenshiftCli() {
	// Método 1: Descargar y extraer .tar.gz
	openshiftcli_tar_url := "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/4.16.9/openshift-client-src.tar.gz"
	openshiftcli_cmd_tar := exec.Command("curl", "-L", openshiftcli_tar_url, "-o", "package.tar.gz")
	err := openshiftcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openshiftcli_zip_url := "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/4.16.9/openshift-client-src.zip"
	openshiftcli_cmd_zip := exec.Command("curl", "-L", openshiftcli_zip_url, "-o", "package.zip")
	err = openshiftcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openshiftcli_bin_url := "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/4.16.9/openshift-client-src.bin"
	openshiftcli_cmd_bin := exec.Command("curl", "-L", openshiftcli_bin_url, "-o", "binary.bin")
	err = openshiftcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openshiftcli_src_url := "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/4.16.9/openshift-client-src.src.tar.gz"
	openshiftcli_cmd_src := exec.Command("curl", "-L", openshiftcli_src_url, "-o", "source.tar.gz")
	err = openshiftcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openshiftcli_cmd_direct := exec.Command("./binary")
	err = openshiftcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
