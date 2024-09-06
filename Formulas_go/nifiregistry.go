package main

// NifiRegistry - Centralized storage & management of NiFi/MiNiFi shared resources
// Homepage: https://nifi.apache.org/projects/registry

import (
	"fmt"
	
	"os/exec"
)

func installNifiRegistry() {
	// Método 1: Descargar y extraer .tar.gz
	nifiregistry_tar_url := "https://www.apache.org/dyn/closer.lua?path=/nifi/1.27.0/nifi-registry-1.27.0-bin.zip"
	nifiregistry_cmd_tar := exec.Command("curl", "-L", nifiregistry_tar_url, "-o", "package.tar.gz")
	err := nifiregistry_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nifiregistry_zip_url := "https://www.apache.org/dyn/closer.lua?path=/nifi/1.27.0/nifi-registry-1.27.0-bin.zip"
	nifiregistry_cmd_zip := exec.Command("curl", "-L", nifiregistry_zip_url, "-o", "package.zip")
	err = nifiregistry_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nifiregistry_bin_url := "https://www.apache.org/dyn/closer.lua?path=/nifi/1.27.0/nifi-registry-1.27.0-bin.zip"
	nifiregistry_cmd_bin := exec.Command("curl", "-L", nifiregistry_bin_url, "-o", "binary.bin")
	err = nifiregistry_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nifiregistry_src_url := "https://www.apache.org/dyn/closer.lua?path=/nifi/1.27.0/nifi-registry-1.27.0-bin.zip"
	nifiregistry_cmd_src := exec.Command("curl", "-L", nifiregistry_src_url, "-o", "source.tar.gz")
	err = nifiregistry_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nifiregistry_cmd_direct := exec.Command("./binary")
	err = nifiregistry_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
