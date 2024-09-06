package main

// Proxify - Portable proxy for capturing, manipulating, and replaying HTTP/HTTPS traffic
// Homepage: https://github.com/projectdiscovery/proxify

import (
	"fmt"
	
	"os/exec"
)

func installProxify() {
	// Método 1: Descargar y extraer .tar.gz
	proxify_tar_url := "https://github.com/projectdiscovery/proxify/archive/refs/tags/v0.0.15.tar.gz"
	proxify_cmd_tar := exec.Command("curl", "-L", proxify_tar_url, "-o", "package.tar.gz")
	err := proxify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proxify_zip_url := "https://github.com/projectdiscovery/proxify/archive/refs/tags/v0.0.15.zip"
	proxify_cmd_zip := exec.Command("curl", "-L", proxify_zip_url, "-o", "package.zip")
	err = proxify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proxify_bin_url := "https://github.com/projectdiscovery/proxify/archive/refs/tags/v0.0.15.bin"
	proxify_cmd_bin := exec.Command("curl", "-L", proxify_bin_url, "-o", "binary.bin")
	err = proxify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proxify_src_url := "https://github.com/projectdiscovery/proxify/archive/refs/tags/v0.0.15.src.tar.gz"
	proxify_cmd_src := exec.Command("curl", "-L", proxify_src_url, "-o", "source.tar.gz")
	err = proxify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proxify_cmd_direct := exec.Command("./binary")
	err = proxify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
