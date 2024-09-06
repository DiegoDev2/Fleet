package main

// Id3v2 - Command-line editor
// Homepage: https://id3v2.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installId3v2() {
	// Método 1: Descargar y extraer .tar.gz
	id3v2_tar_url := "https://downloads.sourceforge.net/project/id3v2/id3v2/0.1.12/id3v2-0.1.12.tar.gz"
	id3v2_cmd_tar := exec.Command("curl", "-L", id3v2_tar_url, "-o", "package.tar.gz")
	err := id3v2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	id3v2_zip_url := "https://downloads.sourceforge.net/project/id3v2/id3v2/0.1.12/id3v2-0.1.12.zip"
	id3v2_cmd_zip := exec.Command("curl", "-L", id3v2_zip_url, "-o", "package.zip")
	err = id3v2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	id3v2_bin_url := "https://downloads.sourceforge.net/project/id3v2/id3v2/0.1.12/id3v2-0.1.12.bin"
	id3v2_cmd_bin := exec.Command("curl", "-L", id3v2_bin_url, "-o", "binary.bin")
	err = id3v2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	id3v2_src_url := "https://downloads.sourceforge.net/project/id3v2/id3v2/0.1.12/id3v2-0.1.12.src.tar.gz"
	id3v2_cmd_src := exec.Command("curl", "-L", id3v2_src_url, "-o", "source.tar.gz")
	err = id3v2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	id3v2_cmd_direct := exec.Command("./binary")
	err = id3v2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: id3lib")
	exec.Command("latte", "install", "id3lib").Run()
}
