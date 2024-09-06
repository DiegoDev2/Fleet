package main

// Graphviz - Graph visualization software from AT&T and Bell Labs
// Homepage: https://graphviz.org/

import (
	"fmt"
	
	"os/exec"
)

func installGraphviz() {
	// Método 1: Descargar y extraer .tar.gz
	graphviz_tar_url := "https://gitlab.com/api/v4/projects/4207231/packages/generic/graphviz-releases/12.1.0/graphviz-12.1.0.tar.xz"
	graphviz_cmd_tar := exec.Command("curl", "-L", graphviz_tar_url, "-o", "package.tar.gz")
	err := graphviz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	graphviz_zip_url := "https://gitlab.com/api/v4/projects/4207231/packages/generic/graphviz-releases/12.1.0/graphviz-12.1.0.tar.xz"
	graphviz_cmd_zip := exec.Command("curl", "-L", graphviz_zip_url, "-o", "package.zip")
	err = graphviz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	graphviz_bin_url := "https://gitlab.com/api/v4/projects/4207231/packages/generic/graphviz-releases/12.1.0/graphviz-12.1.0.tar.xz"
	graphviz_cmd_bin := exec.Command("curl", "-L", graphviz_bin_url, "-o", "binary.bin")
	err = graphviz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	graphviz_src_url := "https://gitlab.com/api/v4/projects/4207231/packages/generic/graphviz-releases/12.1.0/graphviz-12.1.0.tar.xz"
	graphviz_cmd_src := exec.Command("curl", "-L", graphviz_src_url, "-o", "source.tar.gz")
	err = graphviz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	graphviz_cmd_direct := exec.Command("./binary")
	err = graphviz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gd")
exec.Command("latte", "install", "gd")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gts")
exec.Command("latte", "install", "gts")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: librsvg")
exec.Command("latte", "install", "librsvg")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
