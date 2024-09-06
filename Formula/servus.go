package main

// Servus - Library and Utilities for zeroconf networking
// Homepage: https://github.com/HBPVIS/Servus

import (
	"fmt"
	
	"os/exec"
)

func installServus() {
	// Método 1: Descargar y extraer .tar.gz
	servus_tar_url := "https://github.com/HBPVIS/Servus.git"
	servus_cmd_tar := exec.Command("curl", "-L", servus_tar_url, "-o", "package.tar.gz")
	err := servus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	servus_zip_url := "https://github.com/HBPVIS/Servus.git"
	servus_cmd_zip := exec.Command("curl", "-L", servus_zip_url, "-o", "package.zip")
	err = servus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	servus_bin_url := "https://github.com/HBPVIS/Servus.git"
	servus_cmd_bin := exec.Command("curl", "-L", servus_bin_url, "-o", "binary.bin")
	err = servus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	servus_src_url := "https://github.com/HBPVIS/Servus.git"
	servus_cmd_src := exec.Command("curl", "-L", servus_src_url, "-o", "source.tar.gz")
	err = servus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	servus_cmd_direct := exec.Command("./binary")
	err = servus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
}
