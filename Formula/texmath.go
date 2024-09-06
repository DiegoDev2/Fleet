package main

// Texmath - Haskell library for converting LaTeX math to MathML
// Homepage: https://johnmacfarlane.net/texmath.html

import (
	"fmt"
	
	"os/exec"
)

func installTexmath() {
	// Método 1: Descargar y extraer .tar.gz
	texmath_tar_url := "https://hackage.haskell.org/package/texmath-0.12.8.9/texmath-0.12.8.9.tar.gz"
	texmath_cmd_tar := exec.Command("curl", "-L", texmath_tar_url, "-o", "package.tar.gz")
	err := texmath_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	texmath_zip_url := "https://hackage.haskell.org/package/texmath-0.12.8.9/texmath-0.12.8.9.zip"
	texmath_cmd_zip := exec.Command("curl", "-L", texmath_zip_url, "-o", "package.zip")
	err = texmath_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	texmath_bin_url := "https://hackage.haskell.org/package/texmath-0.12.8.9/texmath-0.12.8.9.bin"
	texmath_cmd_bin := exec.Command("curl", "-L", texmath_bin_url, "-o", "binary.bin")
	err = texmath_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	texmath_src_url := "https://hackage.haskell.org/package/texmath-0.12.8.9/texmath-0.12.8.9.src.tar.gz"
	texmath_cmd_src := exec.Command("curl", "-L", texmath_src_url, "-o", "source.tar.gz")
	err = texmath_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	texmath_cmd_direct := exec.Command("./binary")
	err = texmath_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
}
