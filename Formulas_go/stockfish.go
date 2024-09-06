package main

// Stockfish - Strong open-source chess engine
// Homepage: https://stockfishchess.org/

import (
	"fmt"
	
	"os/exec"
)

func installStockfish() {
	// Método 1: Descargar y extraer .tar.gz
	stockfish_tar_url := "https://github.com/official-stockfish/Stockfish/archive/refs/tags/sf_16.1.tar.gz"
	stockfish_cmd_tar := exec.Command("curl", "-L", stockfish_tar_url, "-o", "package.tar.gz")
	err := stockfish_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stockfish_zip_url := "https://github.com/official-stockfish/Stockfish/archive/refs/tags/sf_16.1.zip"
	stockfish_cmd_zip := exec.Command("curl", "-L", stockfish_zip_url, "-o", "package.zip")
	err = stockfish_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stockfish_bin_url := "https://github.com/official-stockfish/Stockfish/archive/refs/tags/sf_16.1.bin"
	stockfish_cmd_bin := exec.Command("curl", "-L", stockfish_bin_url, "-o", "binary.bin")
	err = stockfish_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stockfish_src_url := "https://github.com/official-stockfish/Stockfish/archive/refs/tags/sf_16.1.src.tar.gz"
	stockfish_cmd_src := exec.Command("curl", "-L", stockfish_src_url, "-o", "source.tar.gz")
	err = stockfish_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stockfish_cmd_direct := exec.Command("./binary")
	err = stockfish_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
