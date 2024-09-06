package main

// FernApi - Stripe-level SDKs and Docs for your API
// Homepage: https://buildwithfern.com/

import (
	"fmt"
	
	"os/exec"
)

func installFernApi() {
	// Método 1: Descargar y extraer .tar.gz
	fernapi_tar_url := "https://registry.npmjs.org/fern-api/-/fern-api-0.41.2.tgz"
	fernapi_cmd_tar := exec.Command("curl", "-L", fernapi_tar_url, "-o", "package.tar.gz")
	err := fernapi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fernapi_zip_url := "https://registry.npmjs.org/fern-api/-/fern-api-0.41.2.tgz"
	fernapi_cmd_zip := exec.Command("curl", "-L", fernapi_zip_url, "-o", "package.zip")
	err = fernapi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fernapi_bin_url := "https://registry.npmjs.org/fern-api/-/fern-api-0.41.2.tgz"
	fernapi_cmd_bin := exec.Command("curl", "-L", fernapi_bin_url, "-o", "binary.bin")
	err = fernapi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fernapi_src_url := "https://registry.npmjs.org/fern-api/-/fern-api-0.41.2.tgz"
	fernapi_cmd_src := exec.Command("curl", "-L", fernapi_src_url, "-o", "source.tar.gz")
	err = fernapi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fernapi_cmd_direct := exec.Command("./binary")
	err = fernapi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
