package main

// CloudSqlProxy - Utility for connecting securely to your Cloud SQL instances
// Homepage: https://github.com/GoogleCloudPlatform/cloud-sql-proxy

import (
	"fmt"
	
	"os/exec"
)

func installCloudSqlProxy() {
	// Método 1: Descargar y extraer .tar.gz
	cloudsqlproxy_tar_url := "https://github.com/GoogleCloudPlatform/cloud-sql-proxy/archive/refs/tags/v2.13.0.tar.gz"
	cloudsqlproxy_cmd_tar := exec.Command("curl", "-L", cloudsqlproxy_tar_url, "-o", "package.tar.gz")
	err := cloudsqlproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudsqlproxy_zip_url := "https://github.com/GoogleCloudPlatform/cloud-sql-proxy/archive/refs/tags/v2.13.0.zip"
	cloudsqlproxy_cmd_zip := exec.Command("curl", "-L", cloudsqlproxy_zip_url, "-o", "package.zip")
	err = cloudsqlproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudsqlproxy_bin_url := "https://github.com/GoogleCloudPlatform/cloud-sql-proxy/archive/refs/tags/v2.13.0.bin"
	cloudsqlproxy_cmd_bin := exec.Command("curl", "-L", cloudsqlproxy_bin_url, "-o", "binary.bin")
	err = cloudsqlproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudsqlproxy_src_url := "https://github.com/GoogleCloudPlatform/cloud-sql-proxy/archive/refs/tags/v2.13.0.src.tar.gz"
	cloudsqlproxy_cmd_src := exec.Command("curl", "-L", cloudsqlproxy_src_url, "-o", "source.tar.gz")
	err = cloudsqlproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudsqlproxy_cmd_direct := exec.Command("./binary")
	err = cloudsqlproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
