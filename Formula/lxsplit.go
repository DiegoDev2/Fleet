package main

// Lxsplit - Tool for splitting or joining files
// Homepage: https://lxsplit.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLxsplit() {
	// Método 1: Descargar y extraer .tar.gz
	lxsplit_tar_url := "https://downloads.sourceforge.net/project/lxsplit/lxsplit/0.2.4/lxsplit-0.2.4.tar.gz"
	lxsplit_cmd_tar := exec.Command("curl", "-L", lxsplit_tar_url, "-o", "package.tar.gz")
	err := lxsplit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lxsplit_zip_url := "https://downloads.sourceforge.net/project/lxsplit/lxsplit/0.2.4/lxsplit-0.2.4.zip"
	lxsplit_cmd_zip := exec.Command("curl", "-L", lxsplit_zip_url, "-o", "package.zip")
	err = lxsplit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lxsplit_bin_url := "https://downloads.sourceforge.net/project/lxsplit/lxsplit/0.2.4/lxsplit-0.2.4.bin"
	lxsplit_cmd_bin := exec.Command("curl", "-L", lxsplit_bin_url, "-o", "binary.bin")
	err = lxsplit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lxsplit_src_url := "https://downloads.sourceforge.net/project/lxsplit/lxsplit/0.2.4/lxsplit-0.2.4.src.tar.gz"
	lxsplit_cmd_src := exec.Command("curl", "-L", lxsplit_src_url, "-o", "source.tar.gz")
	err = lxsplit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lxsplit_cmd_direct := exec.Command("./binary")
	err = lxsplit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
