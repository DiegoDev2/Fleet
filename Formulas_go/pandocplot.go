package main

// PandocPlot - Render and include figures in Pandoc documents using many plotting toolkits
// Homepage: https://github.com/LaurentRDC/pandoc-plot

import (
	"fmt"
	
	"os/exec"
)

func installPandocPlot() {
	// Método 1: Descargar y extraer .tar.gz
	pandocplot_tar_url := "https://hackage.haskell.org/package/pandoc-plot-1.8.0/pandoc-plot-1.8.0.tar.gz"
	pandocplot_cmd_tar := exec.Command("curl", "-L", pandocplot_tar_url, "-o", "package.tar.gz")
	err := pandocplot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pandocplot_zip_url := "https://hackage.haskell.org/package/pandoc-plot-1.8.0/pandoc-plot-1.8.0.zip"
	pandocplot_cmd_zip := exec.Command("curl", "-L", pandocplot_zip_url, "-o", "package.zip")
	err = pandocplot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pandocplot_bin_url := "https://hackage.haskell.org/package/pandoc-plot-1.8.0/pandoc-plot-1.8.0.bin"
	pandocplot_cmd_bin := exec.Command("curl", "-L", pandocplot_bin_url, "-o", "binary.bin")
	err = pandocplot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pandocplot_src_url := "https://hackage.haskell.org/package/pandoc-plot-1.8.0/pandoc-plot-1.8.0.src.tar.gz"
	pandocplot_cmd_src := exec.Command("curl", "-L", pandocplot_src_url, "-o", "source.tar.gz")
	err = pandocplot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pandocplot_cmd_direct := exec.Command("./binary")
	err = pandocplot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc")
exec.Command("latte", "install", "ghc")
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
}
