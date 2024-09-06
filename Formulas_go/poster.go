package main

// Poster - Create large posters out of PostScript pages
// Homepage: https://schrfr.github.io/poster/

import (
	"fmt"
	
	"os/exec"
)

func installPoster() {
	// Método 1: Descargar y extraer .tar.gz
	poster_tar_url := "https://github.com/schrfr/poster/archive/refs/tags/1.0.0.tar.gz"
	poster_cmd_tar := exec.Command("curl", "-L", poster_tar_url, "-o", "package.tar.gz")
	err := poster_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	poster_zip_url := "https://github.com/schrfr/poster/archive/refs/tags/1.0.0.zip"
	poster_cmd_zip := exec.Command("curl", "-L", poster_zip_url, "-o", "package.zip")
	err = poster_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	poster_bin_url := "https://github.com/schrfr/poster/archive/refs/tags/1.0.0.bin"
	poster_cmd_bin := exec.Command("curl", "-L", poster_bin_url, "-o", "binary.bin")
	err = poster_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	poster_src_url := "https://github.com/schrfr/poster/archive/refs/tags/1.0.0.src.tar.gz"
	poster_cmd_src := exec.Command("curl", "-L", poster_src_url, "-o", "source.tar.gz")
	err = poster_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	poster_cmd_direct := exec.Command("./binary")
	err = poster_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
