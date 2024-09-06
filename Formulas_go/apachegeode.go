package main

// ApacheGeode - In-memory Data Grid for fast transactional data processing
// Homepage: https://geode.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installApacheGeode() {
	// Método 1: Descargar y extraer .tar.gz
	apachegeode_tar_url := "https://www.apache.org/dyn/closer.lua?path=geode/1.15.1/apache-geode-1.15.1.tgz"
	apachegeode_cmd_tar := exec.Command("curl", "-L", apachegeode_tar_url, "-o", "package.tar.gz")
	err := apachegeode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apachegeode_zip_url := "https://www.apache.org/dyn/closer.lua?path=geode/1.15.1/apache-geode-1.15.1.tgz"
	apachegeode_cmd_zip := exec.Command("curl", "-L", apachegeode_zip_url, "-o", "package.zip")
	err = apachegeode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apachegeode_bin_url := "https://www.apache.org/dyn/closer.lua?path=geode/1.15.1/apache-geode-1.15.1.tgz"
	apachegeode_cmd_bin := exec.Command("curl", "-L", apachegeode_bin_url, "-o", "binary.bin")
	err = apachegeode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apachegeode_src_url := "https://www.apache.org/dyn/closer.lua?path=geode/1.15.1/apache-geode-1.15.1.tgz"
	apachegeode_cmd_src := exec.Command("curl", "-L", apachegeode_src_url, "-o", "source.tar.gz")
	err = apachegeode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apachegeode_cmd_direct := exec.Command("./binary")
	err = apachegeode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
