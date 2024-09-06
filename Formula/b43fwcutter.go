package main

// B43Fwcutter - Extract firmware from Braodcom 43xx driver files
// Homepage: https://wireless.wiki.kernel.org/en/users/drivers/b43

import (
	"fmt"
	
	"os/exec"
)

func installB43Fwcutter() {
	// Método 1: Descargar y extraer .tar.gz
	b43fwcutter_tar_url := "https://bues.ch/b43/fwcutter/b43-fwcutter-019.tar.bz2"
	b43fwcutter_cmd_tar := exec.Command("curl", "-L", b43fwcutter_tar_url, "-o", "package.tar.gz")
	err := b43fwcutter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	b43fwcutter_zip_url := "https://bues.ch/b43/fwcutter/b43-fwcutter-019.tar.bz2"
	b43fwcutter_cmd_zip := exec.Command("curl", "-L", b43fwcutter_zip_url, "-o", "package.zip")
	err = b43fwcutter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	b43fwcutter_bin_url := "https://bues.ch/b43/fwcutter/b43-fwcutter-019.tar.bz2"
	b43fwcutter_cmd_bin := exec.Command("curl", "-L", b43fwcutter_bin_url, "-o", "binary.bin")
	err = b43fwcutter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	b43fwcutter_src_url := "https://bues.ch/b43/fwcutter/b43-fwcutter-019.tar.bz2"
	b43fwcutter_cmd_src := exec.Command("curl", "-L", b43fwcutter_src_url, "-o", "source.tar.gz")
	err = b43fwcutter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	b43fwcutter_cmd_direct := exec.Command("./binary")
	err = b43fwcutter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
