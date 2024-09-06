package main

// Pgpdump - PGP packet visualizer
// Homepage: https://www.mew.org/~kazu/proj/pgpdump/en/

import (
	"fmt"
	
	"os/exec"
)

func installPgpdump() {
	// Método 1: Descargar y extraer .tar.gz
	pgpdump_tar_url := "https://github.com/kazu-yamamoto/pgpdump/archive/refs/tags/v0.36.tar.gz"
	pgpdump_cmd_tar := exec.Command("curl", "-L", pgpdump_tar_url, "-o", "package.tar.gz")
	err := pgpdump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgpdump_zip_url := "https://github.com/kazu-yamamoto/pgpdump/archive/refs/tags/v0.36.zip"
	pgpdump_cmd_zip := exec.Command("curl", "-L", pgpdump_zip_url, "-o", "package.zip")
	err = pgpdump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgpdump_bin_url := "https://github.com/kazu-yamamoto/pgpdump/archive/refs/tags/v0.36.bin"
	pgpdump_cmd_bin := exec.Command("curl", "-L", pgpdump_bin_url, "-o", "binary.bin")
	err = pgpdump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgpdump_src_url := "https://github.com/kazu-yamamoto/pgpdump/archive/refs/tags/v0.36.src.tar.gz"
	pgpdump_cmd_src := exec.Command("curl", "-L", pgpdump_src_url, "-o", "source.tar.gz")
	err = pgpdump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgpdump_cmd_direct := exec.Command("./binary")
	err = pgpdump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
