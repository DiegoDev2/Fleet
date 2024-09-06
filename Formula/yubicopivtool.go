package main

// YubicoPivTool - Command-line tool for the YubiKey PIV application
// Homepage: https://developers.yubico.com/yubico-piv-tool/

import (
	"fmt"
	
	"os/exec"
)

func installYubicoPivTool() {
	// Método 1: Descargar y extraer .tar.gz
	yubicopivtool_tar_url := "https://developers.yubico.com/yubico-piv-tool/Releases/yubico-piv-tool-2.6.0.tar.gz"
	yubicopivtool_cmd_tar := exec.Command("curl", "-L", yubicopivtool_tar_url, "-o", "package.tar.gz")
	err := yubicopivtool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yubicopivtool_zip_url := "https://developers.yubico.com/yubico-piv-tool/Releases/yubico-piv-tool-2.6.0.zip"
	yubicopivtool_cmd_zip := exec.Command("curl", "-L", yubicopivtool_zip_url, "-o", "package.zip")
	err = yubicopivtool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yubicopivtool_bin_url := "https://developers.yubico.com/yubico-piv-tool/Releases/yubico-piv-tool-2.6.0.bin"
	yubicopivtool_cmd_bin := exec.Command("curl", "-L", yubicopivtool_bin_url, "-o", "binary.bin")
	err = yubicopivtool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yubicopivtool_src_url := "https://developers.yubico.com/yubico-piv-tool/Releases/yubico-piv-tool-2.6.0.src.tar.gz"
	yubicopivtool_cmd_src := exec.Command("curl", "-L", yubicopivtool_src_url, "-o", "source.tar.gz")
	err = yubicopivtool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yubicopivtool_cmd_direct := exec.Command("./binary")
	err = yubicopivtool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: check")
	exec.Command("latte", "install", "check").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gengetopt")
	exec.Command("latte", "install", "gengetopt").Run()
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
