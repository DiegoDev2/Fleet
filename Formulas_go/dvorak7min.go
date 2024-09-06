package main

// Dvorak7min - Dvorak (simplified keyboard layout) typing tutor
// Homepage: https://web.archive.org/web/dvorak7min.sourcearchive.com/

import (
	"fmt"
	
	"os/exec"
)

func installDvorak7min() {
	// Método 1: Descargar y extraer .tar.gz
	dvorak7min_tar_url := "https://deb.debian.org/debian/pool/main/d/dvorak7min/dvorak7min_1.6.1+repack.orig.tar.gz"
	dvorak7min_cmd_tar := exec.Command("curl", "-L", dvorak7min_tar_url, "-o", "package.tar.gz")
	err := dvorak7min_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dvorak7min_zip_url := "https://deb.debian.org/debian/pool/main/d/dvorak7min/dvorak7min_1.6.1+repack.orig.zip"
	dvorak7min_cmd_zip := exec.Command("curl", "-L", dvorak7min_zip_url, "-o", "package.zip")
	err = dvorak7min_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dvorak7min_bin_url := "https://deb.debian.org/debian/pool/main/d/dvorak7min/dvorak7min_1.6.1+repack.orig.bin"
	dvorak7min_cmd_bin := exec.Command("curl", "-L", dvorak7min_bin_url, "-o", "binary.bin")
	err = dvorak7min_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dvorak7min_src_url := "https://deb.debian.org/debian/pool/main/d/dvorak7min/dvorak7min_1.6.1+repack.orig.src.tar.gz"
	dvorak7min_cmd_src := exec.Command("curl", "-L", dvorak7min_src_url, "-o", "source.tar.gz")
	err = dvorak7min_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dvorak7min_cmd_direct := exec.Command("./binary")
	err = dvorak7min_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
