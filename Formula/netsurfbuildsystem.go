package main

// NetsurfBuildsystem - Makefiles shared by NetSurf projects
// Homepage: https://source.netsurf-browser.org/buildsystem.git

import (
	"fmt"
	
	"os/exec"
)

func installNetsurfBuildsystem() {
	// Método 1: Descargar y extraer .tar.gz
	netsurfbuildsystem_tar_url := "https://download.netsurf-browser.org/libs/releases/buildsystem-1.10.tar.gz"
	netsurfbuildsystem_cmd_tar := exec.Command("curl", "-L", netsurfbuildsystem_tar_url, "-o", "package.tar.gz")
	err := netsurfbuildsystem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netsurfbuildsystem_zip_url := "https://download.netsurf-browser.org/libs/releases/buildsystem-1.10.zip"
	netsurfbuildsystem_cmd_zip := exec.Command("curl", "-L", netsurfbuildsystem_zip_url, "-o", "package.zip")
	err = netsurfbuildsystem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netsurfbuildsystem_bin_url := "https://download.netsurf-browser.org/libs/releases/buildsystem-1.10.bin"
	netsurfbuildsystem_cmd_bin := exec.Command("curl", "-L", netsurfbuildsystem_bin_url, "-o", "binary.bin")
	err = netsurfbuildsystem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netsurfbuildsystem_src_url := "https://download.netsurf-browser.org/libs/releases/buildsystem-1.10.src.tar.gz"
	netsurfbuildsystem_cmd_src := exec.Command("curl", "-L", netsurfbuildsystem_src_url, "-o", "source.tar.gz")
	err = netsurfbuildsystem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netsurfbuildsystem_cmd_direct := exec.Command("./binary")
	err = netsurfbuildsystem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
