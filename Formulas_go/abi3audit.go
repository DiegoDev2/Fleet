package main

// Abi3audit - Scans Python packages for abi3 violations and inconsistencies
// Homepage: https://github.com/trailofbits/abi3audit

import (
	"fmt"
	
	"os/exec"
)

func installAbi3audit() {
	// Método 1: Descargar y extraer .tar.gz
	abi3audit_tar_url := "https://files.pythonhosted.org/packages/fc/9f/30f21d93b511e250e5fd4987efffc48e3d840048d9a1fbdbb8f70e2f2387/abi3audit-0.0.14.tar.gz"
	abi3audit_cmd_tar := exec.Command("curl", "-L", abi3audit_tar_url, "-o", "package.tar.gz")
	err := abi3audit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abi3audit_zip_url := "https://files.pythonhosted.org/packages/fc/9f/30f21d93b511e250e5fd4987efffc48e3d840048d9a1fbdbb8f70e2f2387/abi3audit-0.0.14.zip"
	abi3audit_cmd_zip := exec.Command("curl", "-L", abi3audit_zip_url, "-o", "package.zip")
	err = abi3audit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abi3audit_bin_url := "https://files.pythonhosted.org/packages/fc/9f/30f21d93b511e250e5fd4987efffc48e3d840048d9a1fbdbb8f70e2f2387/abi3audit-0.0.14.bin"
	abi3audit_cmd_bin := exec.Command("curl", "-L", abi3audit_bin_url, "-o", "binary.bin")
	err = abi3audit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abi3audit_src_url := "https://files.pythonhosted.org/packages/fc/9f/30f21d93b511e250e5fd4987efffc48e3d840048d9a1fbdbb8f70e2f2387/abi3audit-0.0.14.src.tar.gz"
	abi3audit_cmd_src := exec.Command("curl", "-L", abi3audit_src_url, "-o", "source.tar.gz")
	err = abi3audit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abi3audit_cmd_direct := exec.Command("./binary")
	err = abi3audit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
