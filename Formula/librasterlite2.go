package main

// Librasterlite2 - Library to store and retrieve huge raster coverages
// Homepage: https://www.gaia-gis.it/fossil/librasterlite2/index

import (
	"fmt"
	
	"os/exec"
)

func installLibrasterlite2() {
	// Método 1: Descargar y extraer .tar.gz
	librasterlite2_tar_url := "https://www.gaia-gis.it/gaia-sins/librasterlite2-sources/librasterlite2-1.1.0-beta1.tar.gz"
	librasterlite2_cmd_tar := exec.Command("curl", "-L", librasterlite2_tar_url, "-o", "package.tar.gz")
	err := librasterlite2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librasterlite2_zip_url := "https://www.gaia-gis.it/gaia-sins/librasterlite2-sources/librasterlite2-1.1.0-beta1.zip"
	librasterlite2_cmd_zip := exec.Command("curl", "-L", librasterlite2_zip_url, "-o", "package.zip")
	err = librasterlite2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librasterlite2_bin_url := "https://www.gaia-gis.it/gaia-sins/librasterlite2-sources/librasterlite2-1.1.0-beta1.bin"
	librasterlite2_cmd_bin := exec.Command("curl", "-L", librasterlite2_bin_url, "-o", "binary.bin")
	err = librasterlite2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librasterlite2_src_url := "https://www.gaia-gis.it/gaia-sins/librasterlite2-sources/librasterlite2-1.1.0-beta1.src.tar.gz"
	librasterlite2_cmd_src := exec.Command("curl", "-L", librasterlite2_src_url, "-o", "source.tar.gz")
	err = librasterlite2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librasterlite2_cmd_direct := exec.Command("./binary")
	err = librasterlite2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: freexl")
	exec.Command("latte", "install", "freexl").Run()
	fmt.Println("Instalando dependencia: geos")
	exec.Command("latte", "install", "geos").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libgeotiff")
	exec.Command("latte", "install", "libgeotiff").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: librttopo")
	exec.Command("latte", "install", "librttopo").Run()
	fmt.Println("Instalando dependencia: libspatialite")
	exec.Command("latte", "install", "libspatialite").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: minizip")
	exec.Command("latte", "install", "minizip").Run()
	fmt.Println("Instalando dependencia: openjpeg")
	exec.Command("latte", "install", "openjpeg").Run()
	fmt.Println("Instalando dependencia: pixman")
	exec.Command("latte", "install", "pixman").Run()
	fmt.Println("Instalando dependencia: proj")
	exec.Command("latte", "install", "proj").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
