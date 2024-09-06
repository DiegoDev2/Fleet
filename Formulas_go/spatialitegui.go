package main

// SpatialiteGui - GUI tool supporting SpatiaLite
// Homepage: https://www.gaia-gis.it/fossil/spatialite_gui/index

import (
	"fmt"
	
	"os/exec"
)

func installSpatialiteGui() {
	// Método 1: Descargar y extraer .tar.gz
	spatialitegui_tar_url := "https://www.gaia-gis.it/gaia-sins/spatialite-gui-sources/spatialite_gui-2.1.0-beta1.tar.gz"
	spatialitegui_cmd_tar := exec.Command("curl", "-L", spatialitegui_tar_url, "-o", "package.tar.gz")
	err := spatialitegui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spatialitegui_zip_url := "https://www.gaia-gis.it/gaia-sins/spatialite-gui-sources/spatialite_gui-2.1.0-beta1.zip"
	spatialitegui_cmd_zip := exec.Command("curl", "-L", spatialitegui_zip_url, "-o", "package.zip")
	err = spatialitegui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spatialitegui_bin_url := "https://www.gaia-gis.it/gaia-sins/spatialite-gui-sources/spatialite_gui-2.1.0-beta1.bin"
	spatialitegui_cmd_bin := exec.Command("curl", "-L", spatialitegui_bin_url, "-o", "binary.bin")
	err = spatialitegui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spatialitegui_src_url := "https://www.gaia-gis.it/gaia-sins/spatialite-gui-sources/spatialite_gui-2.1.0-beta1.src.tar.gz"
	spatialitegui_cmd_src := exec.Command("curl", "-L", spatialitegui_src_url, "-o", "source.tar.gz")
	err = spatialitegui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spatialitegui_cmd_direct := exec.Command("./binary")
	err = spatialitegui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: freexl")
exec.Command("latte", "install", "freexl")
	fmt.Println("Instalando dependencia: geos")
exec.Command("latte", "install", "geos")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: librasterlite2")
exec.Command("latte", "install", "librasterlite2")
	fmt.Println("Instalando dependencia: librttopo")
exec.Command("latte", "install", "librttopo")
	fmt.Println("Instalando dependencia: libspatialite")
exec.Command("latte", "install", "libspatialite")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: libxlsxwriter")
exec.Command("latte", "install", "libxlsxwriter")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: minizip")
exec.Command("latte", "install", "minizip")
	fmt.Println("Instalando dependencia: openjpeg")
exec.Command("latte", "install", "openjpeg")
	fmt.Println("Instalando dependencia: proj")
exec.Command("latte", "install", "proj")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: virtualpg")
exec.Command("latte", "install", "virtualpg")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: wxwidgets")
exec.Command("latte", "install", "wxwidgets")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
