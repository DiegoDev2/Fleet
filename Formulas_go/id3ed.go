package main

// Id3ed - ID3 tag editor for MP3 files
// Homepage: http://code.fluffytapeworm.com/projects/id3ed

import (
	"fmt"
	
	"os/exec"
)

func installId3ed() {
	// Método 1: Descargar y extraer .tar.gz
	id3ed_tar_url := "http://code.fluffytapeworm.com/projects/id3ed/id3ed-1.10.4.tar.gz"
	id3ed_cmd_tar := exec.Command("curl", "-L", id3ed_tar_url, "-o", "package.tar.gz")
	err := id3ed_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	id3ed_zip_url := "http://code.fluffytapeworm.com/projects/id3ed/id3ed-1.10.4.zip"
	id3ed_cmd_zip := exec.Command("curl", "-L", id3ed_zip_url, "-o", "package.zip")
	err = id3ed_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	id3ed_bin_url := "http://code.fluffytapeworm.com/projects/id3ed/id3ed-1.10.4.bin"
	id3ed_cmd_bin := exec.Command("curl", "-L", id3ed_bin_url, "-o", "binary.bin")
	err = id3ed_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	id3ed_src_url := "http://code.fluffytapeworm.com/projects/id3ed/id3ed-1.10.4.src.tar.gz"
	id3ed_cmd_src := exec.Command("curl", "-L", id3ed_src_url, "-o", "source.tar.gz")
	err = id3ed_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	id3ed_cmd_direct := exec.Command("./binary")
	err = id3ed_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
