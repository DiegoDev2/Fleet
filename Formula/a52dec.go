package main

// A52dec - Library for decoding ATSC A/52 streams (AKA 'AC-3')
// Homepage: https://git.adelielinux.org/community/a52dec/

import (
	"fmt"
	
	"os/exec"
)

func installA52dec() {
	// Método 1: Descargar y extraer .tar.gz
	a52dec_tar_url := "https://distfiles.adelielinux.org/source/a52dec/a52dec-0.8.0.tar.gz"
	a52dec_cmd_tar := exec.Command("curl", "-L", a52dec_tar_url, "-o", "package.tar.gz")
	err := a52dec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	a52dec_zip_url := "https://distfiles.adelielinux.org/source/a52dec/a52dec-0.8.0.zip"
	a52dec_cmd_zip := exec.Command("curl", "-L", a52dec_zip_url, "-o", "package.zip")
	err = a52dec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	a52dec_bin_url := "https://distfiles.adelielinux.org/source/a52dec/a52dec-0.8.0.bin"
	a52dec_cmd_bin := exec.Command("curl", "-L", a52dec_bin_url, "-o", "binary.bin")
	err = a52dec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	a52dec_src_url := "https://distfiles.adelielinux.org/source/a52dec/a52dec-0.8.0.src.tar.gz"
	a52dec_cmd_src := exec.Command("curl", "-L", a52dec_src_url, "-o", "source.tar.gz")
	err = a52dec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	a52dec_cmd_direct := exec.Command("./binary")
	err = a52dec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
