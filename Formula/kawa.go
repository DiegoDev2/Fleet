package main

// Kawa - Programming language for Java (implementation of Scheme)
// Homepage: https://www.gnu.org/software/kawa/

import (
	"fmt"
	
	"os/exec"
)

func installKawa() {
	// Método 1: Descargar y extraer .tar.gz
	kawa_tar_url := "https://ftp.gnu.org/gnu/kawa/kawa-3.1.1.zip"
	kawa_cmd_tar := exec.Command("curl", "-L", kawa_tar_url, "-o", "package.tar.gz")
	err := kawa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kawa_zip_url := "https://ftp.gnu.org/gnu/kawa/kawa-3.1.1.zip"
	kawa_cmd_zip := exec.Command("curl", "-L", kawa_zip_url, "-o", "package.zip")
	err = kawa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kawa_bin_url := "https://ftp.gnu.org/gnu/kawa/kawa-3.1.1.zip"
	kawa_cmd_bin := exec.Command("curl", "-L", kawa_bin_url, "-o", "binary.bin")
	err = kawa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kawa_src_url := "https://ftp.gnu.org/gnu/kawa/kawa-3.1.1.zip"
	kawa_cmd_src := exec.Command("curl", "-L", kawa_src_url, "-o", "source.tar.gz")
	err = kawa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kawa_cmd_direct := exec.Command("./binary")
	err = kawa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
