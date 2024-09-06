package main

// Bittwist - Libcap-based Ethernet packet generator
// Homepage: https://bittwist.sourceforge.io

import (
	"fmt"
	
	"os/exec"
)

func installBittwist() {
	// Método 1: Descargar y extraer .tar.gz
	bittwist_tar_url := "https://downloads.sourceforge.net/project/bittwist/macOS/Bit-Twist%203.8/bittwist-macos-3.8.tar.gz"
	bittwist_cmd_tar := exec.Command("curl", "-L", bittwist_tar_url, "-o", "package.tar.gz")
	err := bittwist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bittwist_zip_url := "https://downloads.sourceforge.net/project/bittwist/macOS/Bit-Twist%203.8/bittwist-macos-3.8.zip"
	bittwist_cmd_zip := exec.Command("curl", "-L", bittwist_zip_url, "-o", "package.zip")
	err = bittwist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bittwist_bin_url := "https://downloads.sourceforge.net/project/bittwist/macOS/Bit-Twist%203.8/bittwist-macos-3.8.bin"
	bittwist_cmd_bin := exec.Command("curl", "-L", bittwist_bin_url, "-o", "binary.bin")
	err = bittwist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bittwist_src_url := "https://downloads.sourceforge.net/project/bittwist/macOS/Bit-Twist%203.8/bittwist-macos-3.8.src.tar.gz"
	bittwist_cmd_src := exec.Command("curl", "-L", bittwist_src_url, "-o", "source.tar.gz")
	err = bittwist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bittwist_cmd_direct := exec.Command("./binary")
	err = bittwist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
