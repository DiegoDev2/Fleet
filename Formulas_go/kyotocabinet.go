package main

// KyotoCabinet - Library of routines for managing a database
// Homepage: https://dbmx.net/kyotocabinet/

import (
	"fmt"
	
	"os/exec"
)

func installKyotoCabinet() {
	// Método 1: Descargar y extraer .tar.gz
	kyotocabinet_tar_url := "https://dbmx.net/kyotocabinet/pkg/kyotocabinet-1.2.80.tar.gz"
	kyotocabinet_cmd_tar := exec.Command("curl", "-L", kyotocabinet_tar_url, "-o", "package.tar.gz")
	err := kyotocabinet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kyotocabinet_zip_url := "https://dbmx.net/kyotocabinet/pkg/kyotocabinet-1.2.80.zip"
	kyotocabinet_cmd_zip := exec.Command("curl", "-L", kyotocabinet_zip_url, "-o", "package.zip")
	err = kyotocabinet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kyotocabinet_bin_url := "https://dbmx.net/kyotocabinet/pkg/kyotocabinet-1.2.80.bin"
	kyotocabinet_cmd_bin := exec.Command("curl", "-L", kyotocabinet_bin_url, "-o", "binary.bin")
	err = kyotocabinet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kyotocabinet_src_url := "https://dbmx.net/kyotocabinet/pkg/kyotocabinet-1.2.80.src.tar.gz"
	kyotocabinet_cmd_src := exec.Command("curl", "-L", kyotocabinet_src_url, "-o", "source.tar.gz")
	err = kyotocabinet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kyotocabinet_cmd_direct := exec.Command("./binary")
	err = kyotocabinet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
