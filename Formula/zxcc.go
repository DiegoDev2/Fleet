package main

// Zxcc - CP/M 2/3 emulator for cross-compiling and CP/M tools under UNIX
// Homepage: https://www.seasip.info/Unix/Zxcc/

import (
	"fmt"
	
	"os/exec"
)

func installZxcc() {
	// Método 1: Descargar y extraer .tar.gz
	zxcc_tar_url := "https://www.seasip.info/Unix/Zxcc/zxcc-0.5.7.tar.gz"
	zxcc_cmd_tar := exec.Command("curl", "-L", zxcc_tar_url, "-o", "package.tar.gz")
	err := zxcc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zxcc_zip_url := "https://www.seasip.info/Unix/Zxcc/zxcc-0.5.7.zip"
	zxcc_cmd_zip := exec.Command("curl", "-L", zxcc_zip_url, "-o", "package.zip")
	err = zxcc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zxcc_bin_url := "https://www.seasip.info/Unix/Zxcc/zxcc-0.5.7.bin"
	zxcc_cmd_bin := exec.Command("curl", "-L", zxcc_bin_url, "-o", "binary.bin")
	err = zxcc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zxcc_src_url := "https://www.seasip.info/Unix/Zxcc/zxcc-0.5.7.src.tar.gz"
	zxcc_cmd_src := exec.Command("curl", "-L", zxcc_src_url, "-o", "source.tar.gz")
	err = zxcc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zxcc_cmd_direct := exec.Command("./binary")
	err = zxcc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
