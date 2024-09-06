package main

// Bench - Command-line benchmark tool
// Homepage: https://github.com/Gabriella439/bench

import (
	"fmt"
	
	"os/exec"
)

func installBench() {
	// Método 1: Descargar y extraer .tar.gz
	bench_tar_url := "https://hackage.haskell.org/package/bench-1.0.13/bench-1.0.13.tar.gz"
	bench_cmd_tar := exec.Command("curl", "-L", bench_tar_url, "-o", "package.tar.gz")
	err := bench_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bench_zip_url := "https://hackage.haskell.org/package/bench-1.0.13/bench-1.0.13.zip"
	bench_cmd_zip := exec.Command("curl", "-L", bench_zip_url, "-o", "package.zip")
	err = bench_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bench_bin_url := "https://hackage.haskell.org/package/bench-1.0.13/bench-1.0.13.bin"
	bench_cmd_bin := exec.Command("curl", "-L", bench_bin_url, "-o", "binary.bin")
	err = bench_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bench_src_url := "https://hackage.haskell.org/package/bench-1.0.13/bench-1.0.13.src.tar.gz"
	bench_cmd_src := exec.Command("curl", "-L", bench_src_url, "-o", "source.tar.gz")
	err = bench_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bench_cmd_direct := exec.Command("./binary")
	err = bench_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.8")
exec.Command("latte", "install", "ghc@9.8")
}
