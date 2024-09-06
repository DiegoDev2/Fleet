package main

// Lcm - Libraries and tools for message passing and data marshalling
// Homepage: https://lcm-proj.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installLcm() {
	// Método 1: Descargar y extraer .tar.gz
	lcm_tar_url := "https://github.com/lcm-proj/lcm/archive/refs/tags/v1.5.0.tar.gz"
	lcm_cmd_tar := exec.Command("curl", "-L", lcm_tar_url, "-o", "package.tar.gz")
	err := lcm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lcm_zip_url := "https://github.com/lcm-proj/lcm/archive/refs/tags/v1.5.0.zip"
	lcm_cmd_zip := exec.Command("curl", "-L", lcm_zip_url, "-o", "package.zip")
	err = lcm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lcm_bin_url := "https://github.com/lcm-proj/lcm/archive/refs/tags/v1.5.0.bin"
	lcm_cmd_bin := exec.Command("curl", "-L", lcm_bin_url, "-o", "binary.bin")
	err = lcm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lcm_src_url := "https://github.com/lcm-proj/lcm/archive/refs/tags/v1.5.0.src.tar.gz"
	lcm_cmd_src := exec.Command("curl", "-L", lcm_src_url, "-o", "source.tar.gz")
	err = lcm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lcm_cmd_direct := exec.Command("./binary")
	err = lcm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
