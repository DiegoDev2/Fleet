package main

// Libwandevent - API for developing event-driven programs
// Homepage: https://web.archive.org/web/20220615162419/https://research.wand.net.nz/software/libwandevent.php

import (
	"fmt"
	
	"os/exec"
)

func installLibwandevent() {
	// Método 1: Descargar y extraer .tar.gz
	libwandevent_tar_url := "https://web.archive.org/web/20220126151045/https://research.wand.net.nz/software/libwandevent/libwandevent-3.0.2.tar.gz"
	libwandevent_cmd_tar := exec.Command("curl", "-L", libwandevent_tar_url, "-o", "package.tar.gz")
	err := libwandevent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libwandevent_zip_url := "https://web.archive.org/web/20220126151045/https://research.wand.net.nz/software/libwandevent/libwandevent-3.0.2.zip"
	libwandevent_cmd_zip := exec.Command("curl", "-L", libwandevent_zip_url, "-o", "package.zip")
	err = libwandevent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libwandevent_bin_url := "https://web.archive.org/web/20220126151045/https://research.wand.net.nz/software/libwandevent/libwandevent-3.0.2.bin"
	libwandevent_cmd_bin := exec.Command("curl", "-L", libwandevent_bin_url, "-o", "binary.bin")
	err = libwandevent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libwandevent_src_url := "https://web.archive.org/web/20220126151045/https://research.wand.net.nz/software/libwandevent/libwandevent-3.0.2.src.tar.gz"
	libwandevent_cmd_src := exec.Command("curl", "-L", libwandevent_src_url, "-o", "source.tar.gz")
	err = libwandevent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libwandevent_cmd_direct := exec.Command("./binary")
	err = libwandevent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
