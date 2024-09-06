package main

// TokyoCabinet - Lightweight database library
// Homepage: https://dbmx.net/tokyocabinet/

import (
	"fmt"
	
	"os/exec"
)

func installTokyoCabinet() {
	// Método 1: Descargar y extraer .tar.gz
	tokyocabinet_tar_url := "https://dbmx.net/tokyocabinet/tokyocabinet-1.4.48.tar.gz"
	tokyocabinet_cmd_tar := exec.Command("curl", "-L", tokyocabinet_tar_url, "-o", "package.tar.gz")
	err := tokyocabinet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tokyocabinet_zip_url := "https://dbmx.net/tokyocabinet/tokyocabinet-1.4.48.zip"
	tokyocabinet_cmd_zip := exec.Command("curl", "-L", tokyocabinet_zip_url, "-o", "package.zip")
	err = tokyocabinet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tokyocabinet_bin_url := "https://dbmx.net/tokyocabinet/tokyocabinet-1.4.48.bin"
	tokyocabinet_cmd_bin := exec.Command("curl", "-L", tokyocabinet_bin_url, "-o", "binary.bin")
	err = tokyocabinet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tokyocabinet_src_url := "https://dbmx.net/tokyocabinet/tokyocabinet-1.4.48.src.tar.gz"
	tokyocabinet_cmd_src := exec.Command("curl", "-L", tokyocabinet_src_url, "-o", "source.tar.gz")
	err = tokyocabinet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tokyocabinet_cmd_direct := exec.Command("./binary")
	err = tokyocabinet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
