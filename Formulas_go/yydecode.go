package main

// Yydecode - Decode yEnc archives
// Homepage: https://yydecode.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installYydecode() {
	// Método 1: Descargar y extraer .tar.gz
	yydecode_tar_url := "https://downloads.sourceforge.net/project/yydecode/yydecode/0.2.10/yydecode-0.2.10.tar.gz"
	yydecode_cmd_tar := exec.Command("curl", "-L", yydecode_tar_url, "-o", "package.tar.gz")
	err := yydecode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yydecode_zip_url := "https://downloads.sourceforge.net/project/yydecode/yydecode/0.2.10/yydecode-0.2.10.zip"
	yydecode_cmd_zip := exec.Command("curl", "-L", yydecode_zip_url, "-o", "package.zip")
	err = yydecode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yydecode_bin_url := "https://downloads.sourceforge.net/project/yydecode/yydecode/0.2.10/yydecode-0.2.10.bin"
	yydecode_cmd_bin := exec.Command("curl", "-L", yydecode_bin_url, "-o", "binary.bin")
	err = yydecode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yydecode_src_url := "https://downloads.sourceforge.net/project/yydecode/yydecode/0.2.10/yydecode-0.2.10.src.tar.gz"
	yydecode_cmd_src := exec.Command("curl", "-L", yydecode_src_url, "-o", "source.tar.gz")
	err = yydecode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yydecode_cmd_direct := exec.Command("./binary")
	err = yydecode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
