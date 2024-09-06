package main

// Mapproxy - Accelerating web map proxy
// Homepage: https://github.com/mapproxy/mapproxy

import (
	"fmt"
	
	"os/exec"
)

func installMapproxy() {
	// Método 1: Descargar y extraer .tar.gz
	mapproxy_tar_url := "https://files.pythonhosted.org/packages/bd/23/7051a8b1226e026df32669c059e3a63a9fd9dbe93ffd2af190f3e6d80825/MapProxy-3.0.1.tar.gz"
	mapproxy_cmd_tar := exec.Command("curl", "-L", mapproxy_tar_url, "-o", "package.tar.gz")
	err := mapproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mapproxy_zip_url := "https://files.pythonhosted.org/packages/bd/23/7051a8b1226e026df32669c059e3a63a9fd9dbe93ffd2af190f3e6d80825/MapProxy-3.0.1.zip"
	mapproxy_cmd_zip := exec.Command("curl", "-L", mapproxy_zip_url, "-o", "package.zip")
	err = mapproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mapproxy_bin_url := "https://files.pythonhosted.org/packages/bd/23/7051a8b1226e026df32669c059e3a63a9fd9dbe93ffd2af190f3e6d80825/MapProxy-3.0.1.bin"
	mapproxy_cmd_bin := exec.Command("curl", "-L", mapproxy_bin_url, "-o", "binary.bin")
	err = mapproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mapproxy_src_url := "https://files.pythonhosted.org/packages/bd/23/7051a8b1226e026df32669c059e3a63a9fd9dbe93ffd2af190f3e6d80825/MapProxy-3.0.1.src.tar.gz"
	mapproxy_cmd_src := exec.Command("curl", "-L", mapproxy_src_url, "-o", "source.tar.gz")
	err = mapproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mapproxy_cmd_direct := exec.Command("./binary")
	err = mapproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: pillow")
	exec.Command("latte", "install", "pillow").Run()
	fmt.Println("Instalando dependencia: proj")
	exec.Command("latte", "install", "proj").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
