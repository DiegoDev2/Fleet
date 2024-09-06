package main

// Gdal - Geospatial Data Abstraction Library
// Homepage: https://www.gdal.org/

import (
	"fmt"
	
	"os/exec"
)

func installGdal() {
	// Método 1: Descargar y extraer .tar.gz
	gdal_tar_url := "https://github.com/OSGeo/gdal/releases/download/v3.9.2/gdal-3.9.2.tar.gz"
	gdal_cmd_tar := exec.Command("curl", "-L", gdal_tar_url, "-o", "package.tar.gz")
	err := gdal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gdal_zip_url := "https://github.com/OSGeo/gdal/releases/download/v3.9.2/gdal-3.9.2.zip"
	gdal_cmd_zip := exec.Command("curl", "-L", gdal_zip_url, "-o", "package.zip")
	err = gdal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gdal_bin_url := "https://github.com/OSGeo/gdal/releases/download/v3.9.2/gdal-3.9.2.bin"
	gdal_cmd_bin := exec.Command("curl", "-L", gdal_bin_url, "-o", "binary.bin")
	err = gdal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gdal_src_url := "https://github.com/OSGeo/gdal/releases/download/v3.9.2/gdal-3.9.2.src.tar.gz"
	gdal_cmd_src := exec.Command("curl", "-L", gdal_src_url, "-o", "source.tar.gz")
	err = gdal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gdal_cmd_direct := exec.Command("./binary")
	err = gdal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: apache-arrow")
	exec.Command("latte", "install", "apache-arrow").Run()
	fmt.Println("Instalando dependencia: cfitsio")
	exec.Command("latte", "install", "cfitsio").Run()
	fmt.Println("Instalando dependencia: epsilon")
	exec.Command("latte", "install", "epsilon").Run()
	fmt.Println("Instalando dependencia: expat")
	exec.Command("latte", "install", "expat").Run()
	fmt.Println("Instalando dependencia: freexl")
	exec.Command("latte", "install", "freexl").Run()
	fmt.Println("Instalando dependencia: geos")
	exec.Command("latte", "install", "geos").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: hdf5")
	exec.Command("latte", "install", "hdf5").Run()
	fmt.Println("Instalando dependencia: imath")
	exec.Command("latte", "install", "imath").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: jpeg-xl")
	exec.Command("latte", "install", "jpeg-xl").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: libaec")
	exec.Command("latte", "install", "libaec").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: libgeotiff")
	exec.Command("latte", "install", "libgeotiff").Run()
	fmt.Println("Instalando dependencia: libheif")
	exec.Command("latte", "install", "libheif").Run()
	fmt.Println("Instalando dependencia: libkml")
	exec.Command("latte", "install", "libkml").Run()
	fmt.Println("Instalando dependencia: liblerc")
	exec.Command("latte", "install", "liblerc").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: libspatialite")
	exec.Command("latte", "install", "libspatialite").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: netcdf")
	exec.Command("latte", "install", "netcdf").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: openexr")
	exec.Command("latte", "install", "openexr").Run()
	fmt.Println("Instalando dependencia: openjpeg")
	exec.Command("latte", "install", "openjpeg").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: poppler")
	exec.Command("latte", "install", "poppler").Run()
	fmt.Println("Instalando dependencia: proj")
	exec.Command("latte", "install", "proj").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: qhull")
	exec.Command("latte", "install", "qhull").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: unixodbc")
	exec.Command("latte", "install", "unixodbc").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
	fmt.Println("Instalando dependencia: xerces-c")
	exec.Command("latte", "install", "xerces-c").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: minizip")
	exec.Command("latte", "install", "minizip").Run()
	fmt.Println("Instalando dependencia: uriparser")
	exec.Command("latte", "install", "uriparser").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
