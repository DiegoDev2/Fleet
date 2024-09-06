package main

// Vnstat - Console-based network traffic monitor
// Homepage: https://humdi.net/vnstat/

import (
	"fmt"
	
	"os/exec"
)

func installVnstat() {
	// Método 1: Descargar y extraer .tar.gz
	vnstat_tar_url := "https://humdi.net/vnstat/vnstat-2.12.tar.gz"
	vnstat_cmd_tar := exec.Command("curl", "-L", vnstat_tar_url, "-o", "package.tar.gz")
	err := vnstat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vnstat_zip_url := "https://humdi.net/vnstat/vnstat-2.12.zip"
	vnstat_cmd_zip := exec.Command("curl", "-L", vnstat_zip_url, "-o", "package.zip")
	err = vnstat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vnstat_bin_url := "https://humdi.net/vnstat/vnstat-2.12.bin"
	vnstat_cmd_bin := exec.Command("curl", "-L", vnstat_bin_url, "-o", "binary.bin")
	err = vnstat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vnstat_src_url := "https://humdi.net/vnstat/vnstat-2.12.src.tar.gz"
	vnstat_cmd_src := exec.Command("curl", "-L", vnstat_src_url, "-o", "source.tar.gz")
	err = vnstat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vnstat_cmd_direct := exec.Command("./binary")
	err = vnstat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gd")
exec.Command("latte", "install", "gd")
}
