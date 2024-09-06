package main

// GraphTool - Efficient network analysis for Python 3
// Homepage: https://graph-tool.skewed.de/

import (
	"fmt"
	
	"os/exec"
)

func installGraphTool() {
	// Método 1: Descargar y extraer .tar.gz
	graphtool_tar_url := "https://downloads.skewed.de/graph-tool/graph-tool-2.77.tar.bz2"
	graphtool_cmd_tar := exec.Command("curl", "-L", graphtool_tar_url, "-o", "package.tar.gz")
	err := graphtool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	graphtool_zip_url := "https://downloads.skewed.de/graph-tool/graph-tool-2.77.tar.bz2"
	graphtool_cmd_zip := exec.Command("curl", "-L", graphtool_zip_url, "-o", "package.zip")
	err = graphtool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	graphtool_bin_url := "https://downloads.skewed.de/graph-tool/graph-tool-2.77.tar.bz2"
	graphtool_cmd_bin := exec.Command("curl", "-L", graphtool_bin_url, "-o", "binary.bin")
	err = graphtool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	graphtool_src_url := "https://downloads.skewed.de/graph-tool/graph-tool-2.77.tar.bz2"
	graphtool_cmd_src := exec.Command("curl", "-L", graphtool_src_url, "-o", "source.tar.gz")
	err = graphtool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	graphtool_cmd_direct := exec.Command("./binary")
	err = graphtool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: boost-python3")
	exec.Command("latte", "install", "boost-python3").Run()
	fmt.Println("Instalando dependencia: cairomm@1.14")
	exec.Command("latte", "install", "cairomm@1.14").Run()
	fmt.Println("Instalando dependencia: cgal")
	exec.Command("latte", "install", "cgal").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: google-sparsehash")
	exec.Command("latte", "install", "google-sparsehash").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: pillow")
	exec.Command("latte", "install", "pillow").Run()
	fmt.Println("Instalando dependencia: py3cairo")
	exec.Command("latte", "install", "py3cairo").Run()
	fmt.Println("Instalando dependencia: pygobject3")
	exec.Command("latte", "install", "pygobject3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: qhull")
	exec.Command("latte", "install", "qhull").Run()
	fmt.Println("Instalando dependencia: scipy")
	exec.Command("latte", "install", "scipy").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: libsigc++@2")
	exec.Command("latte", "install", "libsigc++@2").Run()
	fmt.Println("Instalando dependencia: patchelf")
	exec.Command("latte", "install", "patchelf").Run()
}
