package main

// Graphicsmagick - Image processing tools collection
// Homepage: http://www.graphicsmagick.org/

import (
	"fmt"
	
	"os/exec"
)

func installGraphicsmagick() {
	// Método 1: Descargar y extraer .tar.gz
	graphicsmagick_tar_url := "https://downloads.sourceforge.net/project/graphicsmagick/graphicsmagick/1.3.45/GraphicsMagick-1.3.45.tar.xz"
	graphicsmagick_cmd_tar := exec.Command("curl", "-L", graphicsmagick_tar_url, "-o", "package.tar.gz")
	err := graphicsmagick_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	graphicsmagick_zip_url := "https://downloads.sourceforge.net/project/graphicsmagick/graphicsmagick/1.3.45/GraphicsMagick-1.3.45.tar.xz"
	graphicsmagick_cmd_zip := exec.Command("curl", "-L", graphicsmagick_zip_url, "-o", "package.zip")
	err = graphicsmagick_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	graphicsmagick_bin_url := "https://downloads.sourceforge.net/project/graphicsmagick/graphicsmagick/1.3.45/GraphicsMagick-1.3.45.tar.xz"
	graphicsmagick_cmd_bin := exec.Command("curl", "-L", graphicsmagick_bin_url, "-o", "binary.bin")
	err = graphicsmagick_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	graphicsmagick_src_url := "https://downloads.sourceforge.net/project/graphicsmagick/graphicsmagick/1.3.45/GraphicsMagick-1.3.45.tar.xz"
	graphicsmagick_cmd_src := exec.Command("curl", "-L", graphicsmagick_src_url, "-o", "source.tar.gz")
	err = graphicsmagick_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	graphicsmagick_cmd_direct := exec.Command("./binary")
	err = graphicsmagick_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: jasper")
	exec.Command("latte", "install", "jasper").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: jpeg-xl")
	exec.Command("latte", "install", "jpeg-xl").Run()
	fmt.Println("Instalando dependencia: libheif")
	exec.Command("latte", "install", "libheif").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
