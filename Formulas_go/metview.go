package main

// Metview - Meteorological workstation software
// Homepage: https://metview.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installMetview() {
	// Método 1: Descargar y extraer .tar.gz
	metview_tar_url := "https://confluence.ecmwf.int/download/attachments/51731119/MetviewBundle-2024.4.0-Source.tar.gz"
	metview_cmd_tar := exec.Command("curl", "-L", metview_tar_url, "-o", "package.tar.gz")
	err := metview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	metview_zip_url := "https://confluence.ecmwf.int/download/attachments/51731119/MetviewBundle-2024.4.0-Source.zip"
	metview_cmd_zip := exec.Command("curl", "-L", metview_zip_url, "-o", "package.zip")
	err = metview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	metview_bin_url := "https://confluence.ecmwf.int/download/attachments/51731119/MetviewBundle-2024.4.0-Source.bin"
	metview_cmd_bin := exec.Command("curl", "-L", metview_bin_url, "-o", "binary.bin")
	err = metview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	metview_src_url := "https://confluence.ecmwf.int/download/attachments/51731119/MetviewBundle-2024.4.0-Source.src.tar.gz"
	metview_cmd_src := exec.Command("curl", "-L", metview_src_url, "-o", "source.tar.gz")
	err = metview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	metview_cmd_direct := exec.Command("./binary")
	err = metview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: eccodes")
exec.Command("latte", "install", "eccodes")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: gdbm")
exec.Command("latte", "install", "gdbm")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libaec")
exec.Command("latte", "install", "libaec")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: netcdf")
exec.Command("latte", "install", "netcdf")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: proj")
exec.Command("latte", "install", "proj")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: libtirpc")
exec.Command("latte", "install", "libtirpc")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
}
