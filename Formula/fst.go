package main

// Fst - Represent large sets and maps compactly with finite state transducers
// Homepage: https://github.com/BurntSushi/fst

import (
	"fmt"
	
	"os/exec"
)

func installFst() {
	// Método 1: Descargar y extraer .tar.gz
	fst_tar_url := "https://github.com/BurntSushi/fst/archive/refs/tags/fst-bin-0.4.3.tar.gz"
	fst_cmd_tar := exec.Command("curl", "-L", fst_tar_url, "-o", "package.tar.gz")
	err := fst_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fst_zip_url := "https://github.com/BurntSushi/fst/archive/refs/tags/fst-bin-0.4.3.zip"
	fst_cmd_zip := exec.Command("curl", "-L", fst_zip_url, "-o", "package.zip")
	err = fst_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fst_bin_url := "https://github.com/BurntSushi/fst/archive/refs/tags/fst-bin-0.4.3.bin"
	fst_cmd_bin := exec.Command("curl", "-L", fst_bin_url, "-o", "binary.bin")
	err = fst_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fst_src_url := "https://github.com/BurntSushi/fst/archive/refs/tags/fst-bin-0.4.3.src.tar.gz"
	fst_cmd_src := exec.Command("curl", "-L", fst_src_url, "-o", "source.tar.gz")
	err = fst_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fst_cmd_direct := exec.Command("./binary")
	err = fst_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
