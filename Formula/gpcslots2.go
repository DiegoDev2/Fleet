package main

// Gpcslots2 - Casino text-console game
// Homepage: https://sourceforge.net/projects/gpcslots2/

import (
	"fmt"
	
	"os/exec"
)

func installGpcslots2() {
	// Método 1: Descargar y extraer .tar.gz
	gpcslots2_tar_url := "https://downloads.sourceforge.net/project/gpcslots2/gpcslots2_0-4-5b"
	gpcslots2_cmd_tar := exec.Command("curl", "-L", gpcslots2_tar_url, "-o", "package.tar.gz")
	err := gpcslots2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpcslots2_zip_url := "https://downloads.sourceforge.net/project/gpcslots2/gpcslots2_0-4-5b"
	gpcslots2_cmd_zip := exec.Command("curl", "-L", gpcslots2_zip_url, "-o", "package.zip")
	err = gpcslots2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpcslots2_bin_url := "https://downloads.sourceforge.net/project/gpcslots2/gpcslots2_0-4-5b"
	gpcslots2_cmd_bin := exec.Command("curl", "-L", gpcslots2_bin_url, "-o", "binary.bin")
	err = gpcslots2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpcslots2_src_url := "https://downloads.sourceforge.net/project/gpcslots2/gpcslots2_0-4-5b"
	gpcslots2_cmd_src := exec.Command("curl", "-L", gpcslots2_src_url, "-o", "source.tar.gz")
	err = gpcslots2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpcslots2_cmd_direct := exec.Command("./binary")
	err = gpcslots2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
