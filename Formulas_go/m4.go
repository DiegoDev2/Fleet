package main

// M4 - Macro processing language
// Homepage: https://www.gnu.org/software/m4/

import (
	"fmt"
	
	"os/exec"
)

func installM4() {
	// Método 1: Descargar y extraer .tar.gz
	m4_tar_url := "https://ftp.gnu.org/gnu/m4/m4-1.4.19.tar.xz"
	m4_cmd_tar := exec.Command("curl", "-L", m4_tar_url, "-o", "package.tar.gz")
	err := m4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	m4_zip_url := "https://ftp.gnu.org/gnu/m4/m4-1.4.19.tar.xz"
	m4_cmd_zip := exec.Command("curl", "-L", m4_zip_url, "-o", "package.zip")
	err = m4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	m4_bin_url := "https://ftp.gnu.org/gnu/m4/m4-1.4.19.tar.xz"
	m4_cmd_bin := exec.Command("curl", "-L", m4_bin_url, "-o", "binary.bin")
	err = m4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	m4_src_url := "https://ftp.gnu.org/gnu/m4/m4-1.4.19.tar.xz"
	m4_cmd_src := exec.Command("curl", "-L", m4_src_url, "-o", "source.tar.gz")
	err = m4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	m4_cmd_direct := exec.Command("./binary")
	err = m4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
