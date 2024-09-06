package main

// Fragroute - Intercepts, modifies and rewrites egress traffic for a specified host
// Homepage: https://www.monkey.org/~dugsong/fragroute/

import (
	"fmt"
	
	"os/exec"
)

func installFragroute() {
	// Método 1: Descargar y extraer .tar.gz
	fragroute_tar_url := "https://www.monkey.org/~dugsong/fragroute/fragroute-1.2.tar.gz"
	fragroute_cmd_tar := exec.Command("curl", "-L", fragroute_tar_url, "-o", "package.tar.gz")
	err := fragroute_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fragroute_zip_url := "https://www.monkey.org/~dugsong/fragroute/fragroute-1.2.zip"
	fragroute_cmd_zip := exec.Command("curl", "-L", fragroute_zip_url, "-o", "package.zip")
	err = fragroute_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fragroute_bin_url := "https://www.monkey.org/~dugsong/fragroute/fragroute-1.2.bin"
	fragroute_cmd_bin := exec.Command("curl", "-L", fragroute_bin_url, "-o", "binary.bin")
	err = fragroute_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fragroute_src_url := "https://www.monkey.org/~dugsong/fragroute/fragroute-1.2.src.tar.gz"
	fragroute_cmd_src := exec.Command("curl", "-L", fragroute_src_url, "-o", "source.tar.gz")
	err = fragroute_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fragroute_cmd_direct := exec.Command("./binary")
	err = fragroute_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libdnet")
	exec.Command("latte", "install", "libdnet").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
}
