package main

// GnustepMake - Basic GNUstep Makefiles
// Homepage: https://gnustep.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installGnustepMake() {
	// Método 1: Descargar y extraer .tar.gz
	gnustepmake_tar_url := "https://github.com/gnustep/tools-make/releases/download/make-2_9_2/gnustep-make-2.9.2.tar.gz"
	gnustepmake_cmd_tar := exec.Command("curl", "-L", gnustepmake_tar_url, "-o", "package.tar.gz")
	err := gnustepmake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnustepmake_zip_url := "https://github.com/gnustep/tools-make/releases/download/make-2_9_2/gnustep-make-2.9.2.zip"
	gnustepmake_cmd_zip := exec.Command("curl", "-L", gnustepmake_zip_url, "-o", "package.zip")
	err = gnustepmake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnustepmake_bin_url := "https://github.com/gnustep/tools-make/releases/download/make-2_9_2/gnustep-make-2.9.2.bin"
	gnustepmake_cmd_bin := exec.Command("curl", "-L", gnustepmake_bin_url, "-o", "binary.bin")
	err = gnustepmake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnustepmake_src_url := "https://github.com/gnustep/tools-make/releases/download/make-2_9_2/gnustep-make-2.9.2.src.tar.gz"
	gnustepmake_cmd_src := exec.Command("curl", "-L", gnustepmake_src_url, "-o", "source.tar.gz")
	err = gnustepmake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnustepmake_cmd_direct := exec.Command("./binary")
	err = gnustepmake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
