package main

// Pax - Portable Archive Interchange archive tool
// Homepage: https://www.mirbsd.org/pax.htm

import (
	"fmt"
	
	"os/exec"
)

func installPax() {
	// Método 1: Descargar y extraer .tar.gz
	pax_tar_url := "https://www.mirbsd.org/MirOS/dist/mir/cpio/paxmirabilis-20201030.tgz"
	pax_cmd_tar := exec.Command("curl", "-L", pax_tar_url, "-o", "package.tar.gz")
	err := pax_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pax_zip_url := "https://www.mirbsd.org/MirOS/dist/mir/cpio/paxmirabilis-20201030.tgz"
	pax_cmd_zip := exec.Command("curl", "-L", pax_zip_url, "-o", "package.zip")
	err = pax_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pax_bin_url := "https://www.mirbsd.org/MirOS/dist/mir/cpio/paxmirabilis-20201030.tgz"
	pax_cmd_bin := exec.Command("curl", "-L", pax_bin_url, "-o", "binary.bin")
	err = pax_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pax_src_url := "https://www.mirbsd.org/MirOS/dist/mir/cpio/paxmirabilis-20201030.tgz"
	pax_cmd_src := exec.Command("curl", "-L", pax_src_url, "-o", "source.tar.gz")
	err = pax_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pax_cmd_direct := exec.Command("./binary")
	err = pax_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
