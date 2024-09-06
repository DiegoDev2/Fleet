package main

// Ccls - C/C++/ObjC language server
// Homepage: https://github.com/MaskRay/ccls

import (
	"fmt"
	
	"os/exec"
)

func installCcls() {
	// Método 1: Descargar y extraer .tar.gz
	ccls_tar_url := "https://github.com/MaskRay/ccls/archive/refs/tags/0.20240202.tar.gz"
	ccls_cmd_tar := exec.Command("curl", "-L", ccls_tar_url, "-o", "package.tar.gz")
	err := ccls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ccls_zip_url := "https://github.com/MaskRay/ccls/archive/refs/tags/0.20240202.zip"
	ccls_cmd_zip := exec.Command("curl", "-L", ccls_zip_url, "-o", "package.zip")
	err = ccls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ccls_bin_url := "https://github.com/MaskRay/ccls/archive/refs/tags/0.20240202.bin"
	ccls_cmd_bin := exec.Command("curl", "-L", ccls_bin_url, "-o", "binary.bin")
	err = ccls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ccls_src_url := "https://github.com/MaskRay/ccls/archive/refs/tags/0.20240202.src.tar.gz"
	ccls_cmd_src := exec.Command("curl", "-L", ccls_src_url, "-o", "source.tar.gz")
	err = ccls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ccls_cmd_direct := exec.Command("./binary")
	err = ccls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: rapidjson")
	exec.Command("latte", "install", "rapidjson").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
