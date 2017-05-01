/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import "time"

type StoreResponse struct {
	LastUpdated          time.Time     `json:"last_updated"`
	Channels             []string      `json:"channels"`
	Price                float32       `json:"price"`
	Epoch                string        `json:"epoch"`
	SnapID               string        `json:"snap_id"`
	DeveloperName        string        `json:"developer_name"`
	Confinement          string        `json:"confinement"`
	Release              []string      `json:"release"`
	Plugs                []string      `json:"plugs"`
	Framework            []interface{} `json:"framework"`
	Description          string        `json:"description"`
	TermsOfService       string        `json:"terms_of_service"`
	Channel              string        `json:"channel"`
	Revision             int           `json:"revision"`
	AllowUnauthenticated bool          `json:"allow_unauthenticated"`
	Alias                string        `json:"alias"`
	SnapYamlRaw          string        `json:"snap_yaml_raw"`
	Slots                []string      `json:"slots"`
	Architecture         []string      `json:"architecture"`
	IconUrls             struct {
		Num256 string `json:"256"`
	} `json:"icon_urls"`
	ID              int           `json:"id"`
	AnonDownloadURL string        `json:"anon_download_url"`
	ScreenshotUrls  []interface{} `json:"screenshot_urls"`
	Website         string        `json:"website"`
	Department      []string      `json:"department"`
	Publisher       string        `json:"publisher"`
	Origin          string        `json:"origin"`
	Summary         string        `json:"summary"`
	Promoted        bool          `json:"promoted"`
	Links           struct {
		Curies []struct {
			Templated bool   `json:"templated"`
			Name      string `json:"name"`
			Href      string `json:"href"`
		} `json:"curies"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Changelog      string        `json:"changelog"`
	ScreenshotURL  interface{}   `json:"screenshot_url"`
	Content        string        `json:"content"`
	Version        string        `json:"version"`
	VideoUrls      []interface{} `json:"video_urls"`
	DownloadURL    string        `json:"download_url"`
	RatingsAverage float32       `json:"ratings_average"`
	Prices         struct {
	} `json:"prices"`
	IconURL               string        `json:"icon_url"`
	Title                 string        `json:"title"`
	ClickVersion          string        `json:"click_version"`
	DownloadSha3384       string        `json:"download_sha3_384"`
	DownloadSha512        string        `json:"download_sha512"`
	Status                string        `json:"status"`
	CompanyName           string        `json:"company_name"`
	DatePublished         time.Time     `json:"date_published"`
	License               string        `json:"license"`
	WhitelistCountryCodes []interface{} `json:"whitelist_country_codes"`
	DeveloperID           string        `json:"developer_id"`
	BlacklistCountryCodes []interface{} `json:"blacklist_country_codes"`
	ClickFramework        []interface{} `json:"click_framework"`
	Name                  string        `json:"name"`
	SupportURL            string        `json:"support_url"`
	Keywords              []interface{} `json:"keywords"`
	PackageName           string        `json:"package_name"`
	ReleaseMap            []struct {
		State        string `json:"state"`
		Channel      string `json:"channel"`
		Architecture string `json:"architecture"`
		Series       string `json:"series"`
	} `json:"release_map"`
	BinaryFilesize int `json:"binary_filesize"`
}
