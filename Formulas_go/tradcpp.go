package main

// Tradcpp - K&R-style C preprocessor
// Homepage: https://www.netbsd.org/~dholland/tradcpp/

import (
	"fmt"
	
	"os/exec"
)

func installTradcpp() {
	// Método 1: Descargar y extraer .tar.gz
	tradcpp_tar_url := "https://cdn.netbsd.org/pub/NetBSD/misc/dholland/tradcpp-0.5.3.tar.gz"
	tradcpp_cmd_tar := exec.Command("curl", "-L", tradcpp_tar_url, "-o", "package.tar.gz")
	err := tradcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tradcpp_zip_url := "https://cdn.netbsd.org/pub/NetBSD/misc/dholland/tradcpp-0.5.3.zip"
	tradcpp_cmd_zip := exec.Command("curl", "-L", tradcpp_zip_url, "-o", "package.zip")
	err = tradcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tradcpp_bin_url := "https://cdn.netbsd.org/pub/NetBSD/misc/dholland/tradcpp-0.5.3.bin"
	tradcpp_cmd_bin := exec.Command("curl", "-L", tradcpp_bin_url, "-o", "binary.bin")
	err = tradcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tradcpp_src_url := "https://cdn.netbsd.org/pub/NetBSD/misc/dholland/tradcpp-0.5.3.src.tar.gz"
	tradcpp_cmd_src := exec.Command("curl", "-L", tradcpp_src_url, "-o", "source.tar.gz")
	err = tradcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tradcpp_cmd_direct := exec.Command("./binary")
	err = tradcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bmake")
exec.Command("latte", "install", "bmake")
}
