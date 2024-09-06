package main

// Cocot - Code converter on tty
// Homepage: https://vmi.jp/software/cygwin/cocot.html

import (
	"fmt"
	
	"os/exec"
)

func installCocot() {
	// Método 1: Descargar y extraer .tar.gz
	cocot_tar_url := "https://github.com/vmi/cocot/archive/refs/tags/cocot-1.2-20171118.tar.gz"
	cocot_cmd_tar := exec.Command("curl", "-L", cocot_tar_url, "-o", "package.tar.gz")
	err := cocot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cocot_zip_url := "https://github.com/vmi/cocot/archive/refs/tags/cocot-1.2-20171118.zip"
	cocot_cmd_zip := exec.Command("curl", "-L", cocot_zip_url, "-o", "package.zip")
	err = cocot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cocot_bin_url := "https://github.com/vmi/cocot/archive/refs/tags/cocot-1.2-20171118.bin"
	cocot_cmd_bin := exec.Command("curl", "-L", cocot_bin_url, "-o", "binary.bin")
	err = cocot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cocot_src_url := "https://github.com/vmi/cocot/archive/refs/tags/cocot-1.2-20171118.src.tar.gz"
	cocot_cmd_src := exec.Command("curl", "-L", cocot_src_url, "-o", "source.tar.gz")
	err = cocot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cocot_cmd_direct := exec.Command("./binary")
	err = cocot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
